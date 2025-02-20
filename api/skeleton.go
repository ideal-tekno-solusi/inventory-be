package api

import (
	"app/bootstrap"
	inventory "app/internal/inventory/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, cfg *bootstrap.Container) {
	inventory.RestRegister(r, cfg)
}
