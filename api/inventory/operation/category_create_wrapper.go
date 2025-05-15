package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

		handler(ctx, &params)

		ctx.Next()
	}
}
