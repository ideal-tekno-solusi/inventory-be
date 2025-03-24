package inventory

import (
	"app/api/inventory/operation"
	"app/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, s Service) {
	v1 := r.Group("/v1")
	v1.GET("/api/inventory", operation.InventoryWrapper(s.Inventory))
	v1.GET("/api/category", operation.CategoryWrapper(s.Category))
	v1.POST("/api/category", middleware.CsrfValidatorWrapper(), operation.CategoryCreateWrapper(s.CategoryCreate))
	v1.PATCH("/api/category/:id/update", middleware.CsrfValidatorWrapper(), operation.CategoryUpdateWrapper(s.CategoryUpdate))
	v1.DELETE("/api/category/:id/delete", middleware.CsrfValidatorWrapper(), operation.CategoryDeleteWrapper(s.CategoryDelete))
	v1.GET("/api/login", operation.LoginWrapper(s.Login))
}
