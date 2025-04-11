package operation

import (
	"app/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type CategoryDeleteRequest struct {
	Id string
}

func CategoryDeleteWrapper(handler func(ctx *gin.Context, params *CategoryDeleteRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryDeleteRequest{}

		params.Id = ctx.Param("id")

		err := validateCategoryDeleteReq(params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

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

func validateCategoryDeleteReq(params CategoryDeleteRequest) error {
	if params.Id == "" {
		return errors.New("id can't be empty")
	}

	if len(params.Id) > 20 {
		return errors.New("invalid id length")
	}

	return nil
}
