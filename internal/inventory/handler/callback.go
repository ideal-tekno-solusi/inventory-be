package handler

import (
	"app/api/inventory/operation"
	"app/internal/inventory/entity"
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
	verifierDomain := viper.GetString("config.verifier.domain")
	verifierPath := viper.GetString("config.verifier.path")
	verifierSecure := viper.GetBool("config.verifier.secure")
	verifierHttponly := viper.GetBool("config.verifier.httponly")

	result := entity.TokenResponse{}

	codeVerifier, err := ctx.Cookie("verifier")
	if err != nil || codeVerifier == "" {
		errorMessage := "code verifier not found, please try to login again"
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	//? send http req to /token
	uri := viper.GetString("config.auth.uri")
	path := viper.GetString("config.auth.path.token")

	body := entity.TokenRequest{
		Code:         params.Code,
		CodeVerifier: codeVerifier,
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
	ctx.SetCookie("verifier", "", -1, verifierPath, verifierDomain, verifierSecure, verifierHttponly)

	ctx.JSON(http.StatusOK, result)
}
