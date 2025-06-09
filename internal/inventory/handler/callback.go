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
		errorMessage := fmt.Sprintf("response from server is not ok, response server: %v", string(res))
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, status, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	reqDefaultRes, reqBodyRes, err := utils.BindResponse(res)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to bind response with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	err = json.Unmarshal(*reqBodyRes, &result)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to decode response with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	//? set refresh token cookie
	refreshTokenAge := viper.GetInt("config.refreshToken.age")
	refreshTokenDomain := viper.GetString("config.refreshToken.domain")
	refreshTokenPath := viper.GetString("config.refreshToken.path")
	refreshTokenSecure := viper.GetBool("config.refreshToken.secure")
	refreshTokenHttponly := viper.GetBool("config.refreshToken.httponly")

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("refreshToken", result.RefreshToken, refreshTokenAge, refreshTokenPath, refreshTokenDomain, refreshTokenSecure, refreshTokenHttponly)

	//? clean up
	ctx.SetCookie("verifier", "", -1, verifierPath, verifierDomain, verifierSecure, verifierHttponly)
	ctx.SetCookie("state", "", -1, verifierPath, verifierDomain, verifierSecure, verifierHttponly)

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(reqDefaultRes, true, result))
}
