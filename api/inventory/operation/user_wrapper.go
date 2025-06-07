package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserRequest struct {
	Token string `header:"token" binding:"required"`
}

func UserWrapper(handler func(ctx *gin.Context, params *UserRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := UserRequest{}

		err := ctx.ShouldBindHeader(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusBadRequest, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
