package operation

import "github.com/gin-gonic/gin"

type LoginRequest struct{}

func LoginWrapper(handler func(ctx *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(ctx)

		ctx.Next()
	}
}
