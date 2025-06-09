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

	totalData, totalPage, err := categoryService.CountCategory(ctx, params.Name, params.Limit)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to count categories with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	categories, err := categoryService.FetchCategory(ctx, params.Name, params.Page, params.Limit)
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

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(nil, true, res))
}

func (r *RestService) CategoryCreate(ctx *gin.Context, params *operation.CategoryCreateRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	categoryService := repository.CategoryRepository(repo)

	countExist, _, err := categoryService.CountCategory(ctx, params.Name, 10)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to count categories exist with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	if countExist > 0 {
		errorMessage := fmt.Sprintf("failed to create new category with name %v, name already exist", params.Name)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	err = categoryService.CreateCategory(ctx, params.Name, params.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to create new category with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.Status(http.StatusCreated)
}

func (r *RestService) CategoryUpdate(ctx *gin.Context, params *operation.CategoryUpdateRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	categoryService := repository.CategoryRepository(repo)

	err := categoryService.UpdateCategory(ctx, params.Id, params.Name, params.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to update category with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.Status(http.StatusNoContent)
}

func (r *RestService) CategoryDelete(ctx *gin.Context, params *operation.CategoryDeleteRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	categoryService := repository.CategoryRepository(repo)

	err := categoryService.DeleteCategory(ctx, params.Id)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to soft delete category with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.Status(http.StatusNoContent)
}
