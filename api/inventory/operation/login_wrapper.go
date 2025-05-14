package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type LoginRequest struct {
	RedirectUrl string `query:"redirect-url"`
	CsrfToken   string
}

func LoginWrapper(handler func(ctx *gin.Context, params *LoginRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := LoginRequest{}

		err := ctx.ShouldBindQuery(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusBadRequest, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

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
