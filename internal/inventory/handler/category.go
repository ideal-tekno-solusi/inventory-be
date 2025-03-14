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

func (r *RestService) Category(ctx *gin.Context, params *operation.CategoryRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	categoryService := repository.CategoryRepository(repo)

	var res entity.CategoryResponse

	totalData, totalPage, err := categoryService.CountCategory(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to count categories with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	categories, err := categoryService.FetchCategory(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to fetch categories with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	res = entity.CategoryResponse{
		TotalData:   totalData,
		CurrentPage: params.Page,
		TotalPage:   totalPage,
		Categories:  categories,
	}

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(true, res))
}

func (r *RestService) CategoryCreate(ctx *gin.Context, params *operation.CategoryCreateRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	categoryService := repository.CategoryRepository(repo)

	err := categoryService.CreateCategory(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to create new category with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.Status(http.StatusCreated)
}
