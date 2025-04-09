package handler

import (
	"app/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func (r *RestService) Login(ctx *gin.Context) {
	//TODO: flow generate code verifier + code challenge here
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

	//? set cookie httponly for code verifier, untuk dev di set ke 1 jam
	ctx.SetCookie("INVENTORY-CODE-VERIFIER", string(codeVerifier), 3600, "/", "localhost", false, true)

	//TODO: below create new csrf token, ignore the middleware one and replace it with this new token
	age := viper.GetInt("config.csrf.age")
	domain := viper.GetString("config.csrf.domain")
	path := viper.GetString("config.csrf.path")

	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, err.Error(), ctx.FullPath(), uuid.NewString())

		return
	}

	csrfToken := base64.StdEncoding.EncodeToString(key)

	ctx.SetCookie("INVENTORY-XSRF-TOKEN", csrfToken, age, path, domain, false, false)

	type test struct {
		Csrf      string `json:"csrf"`
		Challenge string `json:"challenge"`
	}

	res := test{
		Csrf:      csrfToken,
		Challenge: codeChallenge,
	}

	//TODO: set redirect to sso be /authorization, ex: /auth?redirect=client/homepage&client_id=client&response_type=code&scopes=profile email&state=csrf-example-client&code_challenge={challenge}&code_challenge_method=S256
	// ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/v1/api/category")
	ctx.JSON(200, res)
}
