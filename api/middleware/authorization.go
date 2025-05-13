package middleware

import (
	"app/utils"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Authorize() gin.HandlerFunc {
	//TODO: need testing
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			errorMessage := fmt.Sprintf("failed to do request because authorization failed with error: %v", "authorization header not found / empty")
			logrus.Warn(errorMessage)

			utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

			return
		}

		tokens := strings.Split(token, " ")

		if tokens[0] != "Bearer" {
			errorMessage := fmt.Sprintf("failed to do request because authorization failed with error: %v", fmt.Sprintf("%v is not a valid token, only accept Bearer token", tokens[0]))
			logrus.Warn(errorMessage)

			utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

			return
		}

		_, err := parseJwt(tokens[1])
		if err != nil {
			errorMessage := fmt.Sprintf("failed to do request because authorization failed with error: %v", err)
			logrus.Warn(errorMessage)

			utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

			return
		}

		ctx.Next()
	}
}

func parseJwt(token string) (jwt.Token, error) {
	pubKeyString := viper.GetString("secret.sso.public")

	block, _ := pem.Decode([]byte(pubKeyString))

	key, err := jwk.ParseKey(block.Bytes, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	tok, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.ECDH_ES(), key), jwt.WithValidate(false))
	if err != nil {
		return nil, err
	}

	return tok, nil
}
