package handler

import (
	"app/api/inventory/operation"
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

func (r *RestService) Login(ctx *gin.Context, params *operation.LoginRequest) {
	//? flow generate code verifier + code challenge here
	verifierAge := viper.GetInt("config.verifier.age")
	verifierDomain := viper.GetString("config.verifier.domain")
	verifierPath := viper.GetString("config.verifier.path")
	verifierSecure := viper.GetBool("config.verifier.secure")
	verifierHttponly := viper.GetBool("config.verifier.httponly")

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

	urlLogin := viper.GetString("config.url.redirect_fe.login")

	//? set code verifier to cookie
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("verifier", codeVerifierString, verifierAge, verifierPath, verifierDomain, verifierSecure, verifierHttponly)

	redParams := url.Values{}
	redParams.Add("responseType", "code")
	redParams.Add("clientId", "inventory")
	redParams.Add("redirectUrl", params.RedirectUrl)
	redParams.Add("scopes", "user inventory")
	redParams.Add("state", uuid.NewString())
	redParams.Add("codeChallenge", codeChallenge)
	redParams.Add("codeChallengeMethod", "S256")

	//TODO: redirect ke frontend dengan query yg udah dibikin diatas, nanti fe ambil semua query dari redirect dan kirim ke sso /login (body dan query POST) baru ke sso /authorize
	ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%v?%v", urlLogin, redParams.Encode()))
}
