package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type CategoryCreateRequest struct {
	Name        string `json:"name" binding:"required,max=255"`
	Description string `json:"description" binding:"required"`
}

func CategoryCreateWrapper(handler func(ctx *gin.Context, params *CategoryCreateRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryCreateRequest{}

		err := ctx.ShouldBindBodyWithJSON(&params)
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

		handler(ctx, &params)

		ctx.Next()
	}
}
