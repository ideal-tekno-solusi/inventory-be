package middleware

import (
	"app/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Authorize() gin.HandlerFunc {
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

		//TODO: lanjutin validasi tokens[1] yg isinya jwt apakah valid menggunakan private key punya SSO, liat di config.yaml api sso untuk valuenya
		ctx.Next()
	}
}
