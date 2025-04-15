package operation

import (
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
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

		csrfToken := csrf.Token(ctx.Request)
		if csrfToken == "" {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, "csrf token is empty, please try again", ctx.FullPath(), uuid.NewString())

			return
		}

		ctx.Header("X-CSRF-Token", csrfToken)

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
