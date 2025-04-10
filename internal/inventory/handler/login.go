package handler

import (
	"app/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	csrfToken, csrfTokenExist := ctx.Get("INVENTORY-XSRF-TOKEN")
	if !csrfTokenExist {
		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, "csrf token not found, please try again or contact our admin", ctx.FullPath(), uuid.NewString())

		return
	}

	type test struct {
		Csrf      string `json:"csrf"`
		Challenge string `json:"challenge"`
	}

	res := test{
		Csrf:      csrfToken.(string),
		Challenge: codeChallenge,
	}

	//TODO: set redirect to sso be /authorization, ex: /auth?redirect=client/homepage&client_id=client&response_type=code&scopes=profile email&state=csrf-example-client&code_challenge={challenge}&code_challenge_method=S256
	// ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/v1/api/category")
	ctx.JSON(200, res)
}
