package middleware

import (
	"app/utils"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type CsrfValidatorRequest struct {
	XXsrfToken string `json:"INVENTORY-XSRF-TOKEN"`
}

func CsrfGenerateWrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		csrfTokenCookie, _ := ctx.Cookie("INVENTORY-XSRF-TOKEN")

		if csrfTokenCookie != "" {
			ctx.Set("INVENTORY-XSRF-TOKEN", csrfTokenCookie)

			ctx.Next()
		}

		age := viper.GetInt("config.csrf.age")
		domain := viper.GetString("config.csrf.domain")
		path := viper.GetString("config.csrf.path")

		key := make([]byte, 32)
		_, err := rand.Read(key)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		encode := base64.StdEncoding.EncodeToString(key)

		ctx.SetCookie("INVENTORY-XSRF-TOKEN", encode, age, path, domain, false, false)
		ctx.Set("INVENTORY-XSRF-TOKEN", encode)

		ctx.Next()
	}
}

func CsrfValidatorWrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		csrfTokenCookie, _ := ctx.Cookie("INVENTORY-XSRF-TOKEN")

		if csrfTokenCookie == "" {
			utils.SendProblemDetailJson(ctx, http.StatusForbidden, "csrf token is empty", ctx.FullPath(), uuid.NewString())

			ctx.Abort()
			return
		}

		params := CsrfValidatorRequest{}

		err := ctx.ShouldBindBodyWithJSON(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			ctx.Abort()
			return
		}

		err = validateCsrfValidator(params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

			ctx.Abort()
			return
		}

		csrfToken, _ := url.QueryUnescape(params.XXsrfToken)

		if csrfTokenCookie != csrfToken {
			utils.SendProblemDetailJson(ctx, http.StatusForbidden, "csrf token is not match", ctx.FullPath(), uuid.NewString())

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func validateCsrfValidator(params CsrfValidatorRequest) error {
	if params.XXsrfToken == "" {
		return errors.New("csrf token request is empty")
	}

	return nil
}
