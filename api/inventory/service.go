package inventory

import (
	"app/api/inventory/operation"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Inventory(ctx *gin.Context, params *operation.InventoryRequest)
	Category(ctx *gin.Context, params *operation.CategoryRequest)
	CategoryCreate(ctx *gin.Context, params *operation.CategoryCreateRequest)
}
