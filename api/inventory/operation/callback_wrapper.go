package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CallbackRequest struct {
	Code  string `form:"code" binding:"required"`
	State string `form:"state" binding:"required"`
}

func CallbackWrapper(handler func(ctx *gin.Context, params *CallbackRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CallbackRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusInternalServerError, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		if verifyState(ctx, params) {
			utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, "state not found or invalid, please try login again", ctx.FullPath(), uuid.NewString())

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}

func verifyState(ctx *gin.Context, params CallbackRequest) bool {
	state, err := ctx.Cookie("verifier")
	if err != nil || state == "" {
		return false
	}

	if params.State != state {
		return false
	}

	return true
}
