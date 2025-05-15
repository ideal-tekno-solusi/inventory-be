package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryDeleteRequest struct {
	Id string `uri:"id" binding:"required,max=20"`
}

func CategoryDeleteWrapper(handler func(ctx *gin.Context, params *CategoryDeleteRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryDeleteRequest{}

		err := ctx.ShouldBindUri(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusBadRequest, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
