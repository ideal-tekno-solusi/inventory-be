package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CallbackRequest struct {
	Code string `form:"code" binding:"required"`
}

func CallbackWrapper(handler func(ctx *gin.Context, params *CallbackRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CallbackRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusInternalServerError, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
