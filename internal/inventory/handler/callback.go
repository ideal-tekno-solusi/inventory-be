package handler

import (
	"app/api/inventory/operation"
	"app/internal/inventory/entity"
	"app/internal/inventory/repository"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (r *RestService) Callback(ctx *gin.Context, params *operation.CallbackRequest) {
	repo := repository.InitRepo(r.dbr, r.dbw)
	callbackService := repository.CallbackRepository(repo)

	message := entity.CodeMessage{}
	result := entity.TokenResponse{}

	text, err := utils.DecryptJwe(params.Code)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to decrypt message with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	err = json.Unmarshal([]byte(*text), &message)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to unmarshal message with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	data, err := callbackService.GetChallenge(ctx, message.CodeChallenge)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to get challenge with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	//? send http req to /token
	uri := viper.GetString("config.auth.uri")
	path := viper.GetString("config.auth.path.token")

	body := entity.TokenRequest{
		Code:         message.AuthorizationCode,
		CodeVerifier: data.CodeVerifier.String,
	}

	bodyString, _ := json.Marshal(body)

	status, res, err := utils.SendHttpPostRequest(fmt.Sprintf("%v%v", uri, path), bodyString, nil)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to req token with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}
	if status != http.StatusOK {
		errorMessage := fmt.Sprintf("response from server is not ok, status %v", status)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, status, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to decode response with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	//? clean up
	err = callbackService.DeleteChallenge(ctx, message.CodeChallenge)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to delete challenge with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.JSON(http.StatusOK, result)
}
