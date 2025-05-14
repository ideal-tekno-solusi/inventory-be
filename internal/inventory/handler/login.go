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
)

func (r *RestService) Login(ctx *gin.Context, params *operation.LoginRequest) {
	//? flow generate code verifier + code challenge here
	codeVerifier := make([]byte, 128)
	_, err := rand.Read(codeVerifier)
	if err != nil {
		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

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

	redParams := url.Values{}
	redParams.Add("response_type", "code")
	redParams.Add("client_id", "inventory")
	redParams.Add("redirect_url", params.RedirectUrl)
	redParams.Add("scopes", "user inventory")
	redParams.Add("state", params.CsrfToken)
	redParams.Add("code_challenge", codeChallenge)
	redParams.Add("code_challenge_method", "S256")

	ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://localhost:8081/v1/api/authorization?%v", redParams.Encode()))
}
