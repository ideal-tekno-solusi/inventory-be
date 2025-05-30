package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type InventoryRequest struct {
	Category string `form:"category"`
	BranchId string `form:"branchId"`
	Page     int    `form:"page" binding:"gt=0"`
	Limit    int    `form:"limit" binding:"gte=10,lte=50"`
}

func InventoryWrapper(handler func(ctx *gin.Context, params *InventoryRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := InventoryRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJsonValidate(ctx, http.StatusInternalServerError, "validation error", ctx.FullPath(), uuid.NewString(), err.(validator.ValidationErrors))

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}
