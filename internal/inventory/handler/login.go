package handler

import (
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

func (r *RestService) Login(ctx *gin.Context) {
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

	csrfToken, csrfTokenExist := ctx.Get("INVENTORY-XSRF-TOKEN")
	if !csrfTokenExist {
		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, "csrf token not found, please try again or contact our admin", ctx.FullPath(), uuid.NewString())

		return
	}

	params := url.Values{}
	params.Add("redirect_url", "https://google.com")
	params.Add("client_id", "inventory")
	params.Add("response_type", "code")
	params.Add("scopes", "user inventory")
	params.Add("state", csrfToken.(string))
	params.Add("code_challenge", codeChallenge)
	params.Add("code_challenge_method", "S256")

	ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://localhost:8081/v1/api/authorization?%v", params.Encode()))
}
