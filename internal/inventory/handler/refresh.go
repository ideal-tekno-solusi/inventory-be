package handler

import (
	"app/api/inventory/operation"
	"app/internal/inventory/entity"
	"app/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (r *RestService) RefreshToken(ctx *gin.Context, params *operation.RefreshTokenRequest) {
	//? flow generate code verifier + code challenge here
	verifierAge := viper.GetInt("config.verifier.age")
	verifierDomain := viper.GetString("config.verifier.domain")
	verifierPath := viper.GetString("config.verifier.path")
	verifierSecure := viper.GetBool("config.verifier.secure")
	verifierHttponly := viper.GetBool("config.verifier.httponly")
	redirectUrl := viper.GetString("config.url.default.home")
	authorizeRes := entity.AuthorizeResponse{}
	response := entity.RefreshResponse{}

	if params.RedirectUrl != "" {
		redirectUrl = params.RedirectUrl
	}

	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		errorMessage := "refresh token not found, please try login again"
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	codeVerifier := make([]byte, 128)
	_, err = rand.Read(codeVerifier)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to generate code verifier with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	codeVerifierString := base64.URLEncoding.EncodeToString(codeVerifier)

	hash := sha256.New()

	hash.Write([]byte(codeVerifierString))
	codeChallenge := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	uriAuth := viper.GetString("config.auth.uri")
	pathAuth := viper.GetString("config.auth.path.auth")

	//? set code verifier to cookie
	ctx.SetCookie("verifier", codeVerifierString, verifierAge, verifierPath, verifierDomain, verifierSecure, verifierHttponly)

	query := url.Values{}
	query.Add("responseType", "refresh")
	query.Add("clientId", "inventory")
	query.Add("redirectUrl", redirectUrl)
	query.Add("scopes", "user inventory")
	query.Add("state", refreshToken)
	query.Add("codeChallenge", codeChallenge)
	query.Add("codeChallengeMethod", "S256")

	//? call authorize
	status, res, err := utils.SendHttpGetRequest(fmt.Sprintf("%v%v", uriAuth, pathAuth), &query, nil)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to request authorize with error: %v", err)
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

	err = json.Unmarshal(*reqBodyRes, &authorizeRes)
	if err != nil {
		errorMessage := fmt.Sprintf("failed to decode response with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	callbackUrl := viper.GetString("config.url.default.callbackUrl")
	redParams := url.Values{}
	redParams.Add("code", authorizeRes.AuthorizeCode)
	redParams.Add("state", refreshToken)

	response.CallbackUrl = fmt.Sprintf("%v?%v", callbackUrl, redParams.Encode())

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(reqDefaultRes, true, response))
}
