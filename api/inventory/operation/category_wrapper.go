package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name  string `form:"name"`
	Page  int    `form:"page" binding:"gt=0"`
	Limit int    `form:"limit" binding:"gte=10,lte=50"`
}

func CategoryWrapper(handler func(ctx *gin.Context, params *CategoryRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusInternalServerError, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
