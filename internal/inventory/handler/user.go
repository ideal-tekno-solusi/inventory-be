package handler

import (
	"app/api/inventory/operation"
	"app/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *RestService) User(ctx *gin.Context, params *operation.UserRequest) {
	token := strings.Split(params.Token, " ")

	plainText, err := utils.DecryptJwt(token[1])
	if err != nil {
		errorMessage := fmt.Sprintf("failed to decrypt jwt with error: %v", err)
		logrus.Error(errorMessage)

		utils.SendProblemDetailJson(ctx, http.StatusInternalServerError, errorMessage, ctx.FullPath(), uuid.NewString())

		return
	}

	ctx.JSON(http.StatusOK, utils.GenerateResponseJson(true, plainText))
}
