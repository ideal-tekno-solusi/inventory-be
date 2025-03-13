package operation

import (
	"app/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InventoryRequest struct {
	Category   string `form:"category"`
	LocationId string `form:"locationId"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
}

func InventoryWrapper(handler func(ctx *gin.Context, params *InventoryRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := InventoryRequest{}

		err := ctx.Bind(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		params = defaultValueInventory(params)

		err = validateInventoryReq(params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}

func validateInventoryReq(params InventoryRequest) error {
	if params.Page < 0 {
		return errors.New("page value can't lower than 1")
	}

	if params.Limit < 0 {
		return errors.New("limit value can't lower than 10")
	}

	return nil
}

func defaultValueInventory(params InventoryRequest) InventoryRequest {
	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	return params
}
