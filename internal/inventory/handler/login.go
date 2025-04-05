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

	//TODO: set redirect to sso be /authorization, ex: /auth?redirect=client/homepage&client_id=client&response_type=code&scopes=profile email&state=csrf-example-client&code_challenge={challenge}&code_challenge_method=S256
	ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8080/v1/api/category")
}
