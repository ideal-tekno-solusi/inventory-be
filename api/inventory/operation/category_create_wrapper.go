package operation

import (
	"app/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CategoryCreateWrapper(handler func(ctx *gin.Context, params *CategoryCreateRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryCreateRequest{}

		err := ctx.BindJSON(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		err = validateCategoryCreateReq(params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}

func validateCategoryCreateReq(params CategoryCreateRequest) error {
	if params.Name == "" {
		return errors.New("name can't be empty")
	}

	if len(params.Name) > 255 {
		return errors.New("name max length is 255 character")
	}

	if params.Description == "" {
		return errors.New("description can't be empty")
	}

	return nil
}
