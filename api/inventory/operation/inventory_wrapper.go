package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InventoryRequest struct {
	Category   string `form:"category"`
	LocationId string `form:"locationId"`
}

func InventoryWrapper(handler func(ctx *gin.Context, params *InventoryRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := InventoryRequest{}

		err := ctx.Bind(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		// err = validateInventoryReq(params)
		// if err != nil {
		// 	utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

		// 	return
		// }

		handler(ctx, &params)

		ctx.Next()
	}
}

// func validateInventoryReq(params InventoryRequest) error {
// 	if params.Category == "" {
// 		return errors.New("category can't be empty")
// 	}

// 	if params.LocationId == "" {
// 		return errors.New("location id can't be empty")
// 	}

// 	return nil
// }
