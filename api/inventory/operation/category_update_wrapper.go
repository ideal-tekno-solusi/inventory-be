package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type CategoryUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,max=20"`
}

type CategoryUpdateRequest struct {
	Id          string
	Name        string `json:"name" binding:"required,max=255"`
	Description string `json:"description" binding:"required"`
}

func CategoryUpdateWrapper(handler func(ctx *gin.Context, params *CategoryUpdateRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uriParams := CategoryUpdateRequestUri{}
		params := CategoryUpdateRequest{}

		err := ctx.ShouldBindUri(&uriParams)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusBadRequest, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		err = ctx.ShouldBindBodyWithJSON(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		params.Id = uriParams.Id

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
