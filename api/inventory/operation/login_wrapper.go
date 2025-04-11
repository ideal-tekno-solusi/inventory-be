package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type LoginRequest struct {
	CsrfToken string
}

func LoginWrapper(handler func(ctx *gin.Context, params *LoginRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := LoginRequest{}

		csrfToken := csrf.Token(ctx.Request)
		if csrfToken == "" {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, "csrf token is empty, please try again", ctx.FullPath(), uuid.NewString())

			return
		}

		ctx.Header("X-CSRF-Token", csrfToken)

		params.CsrfToken = csrfToken

		handler(ctx, &params)

		ctx.Next()
	}
}
