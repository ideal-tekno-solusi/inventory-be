package operation

import (
	"app/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryUpdateRequest struct {
	Id          string
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CategoryUpdateWrapper(handler func(ctx *gin.Context, params *CategoryUpdateRequest)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := CategoryUpdateRequest{}

		params.Id = ctx.Param("id")

		err := ctx.ShouldBindBodyWithJSON(&params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		err = validateCategoryUpdateReq(params)
		if err != nil {
			utils.SendProblemDetailJson(ctx, http.StatusBadRequest, err.Error(), ctx.FullPath(), uuid.NewString())

			return
		}

		handler(ctx, &params)

		ctx.Next()
	}
}

func validateCategoryUpdateReq(params CategoryUpdateRequest) error {
	if params.Id == "" {
		return errors.New("id can't be empty")
	}

	if len(params.Id) > 20 {
		return errors.New("invalid id length")
	}

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
