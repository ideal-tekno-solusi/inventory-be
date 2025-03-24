package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *RestService) Login(ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/v1/api/category")
}
