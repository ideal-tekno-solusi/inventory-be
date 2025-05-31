package inventory

import (
	"app/api/inventory/operation"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, s Service) {
	v1 := r.Group("/v1")
	v1.GET("/api/inventory", operation.InventoryWrapper(s.Inventory))
	v1.GET("/api/category", operation.CategoryWrapper(s.Category))
	v1.POST("/api/category", operation.CategoryCreateWrapper(s.CategoryCreate))
	v1.PATCH("/api/category/:id/update", operation.CategoryUpdateWrapper(s.CategoryUpdate))
	v1.DELETE("/api/category/:id/delete", operation.CategoryDeleteWrapper(s.CategoryDelete))
	v1.GET("/api/login", operation.LoginWrapper(s.Login))
	v1.GET("/api/callback", operation.CallbackWrapper(s.Callback))
	v1.GET("/api/token/refresh", operation.RefreshTokenWrapper(s.RefreshToken))
}
