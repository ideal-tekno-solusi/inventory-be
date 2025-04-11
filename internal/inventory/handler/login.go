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
	"github.com/spf13/viper"
)

func (r *RestService) Login(ctx *gin.Context, params *operation.LoginRequest) {
	//? flow generate code verifier + code challenge here
	//TODO: this codeVerifier need to be improve so that the string can contain special char
	codeVerifier := make([]byte, 128)
	_, err := rand.Read(codeVerifier)
	if err != nil {
		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

		return
	}

	hash := sha256.New()

	hash.Write(codeVerifier)
	codeChallenge := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	age := viper.GetInt("config.oauth.age")
	domain := viper.GetString("config.oauth.domain")
	path := viper.GetString("config.oauth.path")

	//? set cookie httponly for code verifier
	ctx.SetCookie("INVENTORY-CODE-VERIFIER", string(codeVerifier), age, path, domain, false, true)

	redParams := url.Values{}
	//TODO: change after dev, preferable to set redirect url from req
	redParams.Add("redirect_url", "https://google.com")
	redParams.Add("client_id", "inventory")
	redParams.Add("response_type", "code")
	redParams.Add("scopes", "user inventory")
	redParams.Add("state", params.CsrfToken)
	redParams.Add("code_challenge", codeChallenge)
	redParams.Add("code_challenge_method", "S256")

	ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://localhost:8081/v1/api/authorization?%v", redParams.Encode()))
}
