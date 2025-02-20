package handler

import (
	"app/api/inventory/operation"
	database "app/database/main"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

func (r *RestService) Inventory(ctx *gin.Context, params *operation.InventoryRequest) {
	queries := database.New(r.dbr)

	args := database.FetchInventoryItemsParams{
		CategoryID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.Category),
			Valid:  true,
		},
		LocationID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.LocationId),
			Valid:  true,
		},
	}

	items, err := queries.FetchInventoryItems(ctx, args)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to fetch items with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(true, items))
}
