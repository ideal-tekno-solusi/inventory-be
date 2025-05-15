package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name  string `form:"name"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

func CategoryWrapper(handler func(ctx *gin.Context, params *CategoryRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryRequest{}

		err := ctx.ShouldBind(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		params = defaultValueCategory(params)

		handler(ctx, &params)

		ctx.Next()
	}
}

func defaultValueCategory(params CategoryRequest) CategoryRequest {
	if params.Page <= 0 {
		params.Page = 1
	}

	if params.Limit <= 10 {
		params.Limit = 10
	}

	return params
}
