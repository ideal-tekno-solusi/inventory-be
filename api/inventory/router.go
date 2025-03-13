package inventory

import (
	"app/api/inventory/operation"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, s Service) {
	r.GET("/inventory", operation.InventoryWrapper(s.Inventory))
	r.GET("/category", operation.CategoryWrapper(s.Category))
}
