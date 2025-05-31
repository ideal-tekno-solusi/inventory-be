package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RefreshTokenRequest struct{}

func RefreshTokenWrapper(handler func(ctx *gin.Context, params *RefreshTokenRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := RefreshTokenRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusBadRequest, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
