package handler

import (
	"app/api/inventory/operation"
	"app/internal/inventory/entity"
	"app/internal/inventory/repository"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *RestService) Inventory(ctx *gin.Context, params *operation.InventoryRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	inventoryService := repository.InventoryRepository(repo)

	var res entity.InventoryResponse

	total, page, err := inventoryService.CountInventoryItems(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to count inventory item with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	items, err := inventoryService.FetchInventoryItems(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to fetch items with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	res = entity.InventoryResponse{
		TotalData:   total,
		CurrentPage: params.Page,
		TotalPage:   page,
		Items:       items,
	}

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(nil, true, res))
}
