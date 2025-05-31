package handler

import (
	"app/api/inventory/operation"
	"app/internal/inventory/repository"
	"app/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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
	refreshToken := ctx.GetHeader("refresh-token")
	if refreshToken == "" {
		errorMessage := "refresh token not found, please try login again"
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusUnauthorized, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	codeVerifier := make([]byte, 128)
	_, err := rand.Read(codeVerifier)
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

	repo := repository.InitRepo(r.dbr, r.dbw)
	loginService := repository.LoginRepository(repo)

	err = loginService.CreateChallenge(ctx, codeVerifierString, codeChallenge, "S256")
	if err != nil {
		errorMessage := fmt.Sprintf("failed to create new challenge with error: %v", err)
		logrus.Warn(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	uriAuth := viper.GetString("config.auth.uri")
	pathAuth := viper.GetString("config.auth.path.auth")

	//TODO: sementara redirect nya example aja, klo diliat mah di flow login juga ini redirect ga kepake sih
	query := url.Values{}
	query.Add("response_type", "refresh")
	query.Add("client_id", "inventory")
	query.Add("redirect_url", "http://example.com")
	query.Add("scopes", "user inventory")
	query.Add("state", refreshToken)
	query.Add("code_challenge", codeChallenge)
	query.Add("code_challenge_method", "S256")

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("%v%v?%v", uriAuth, pathAuth, query.Encode()))
}
