package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Problem struct {
	Type     string      `json:"type"`
	Title    string      `json:"title"`
	Status   int         `json:"status"`
	Detail   interface{} `json:"detail,omitempty"`
	Message  string      `json:"message"`
	Errors   interface{} `json:"errors,omitempty"`
	Instance string      `json:"instance"`
	Guid     string      `json:"guid"`
}

func generateProblemJson(statusCode int, message, instance, guid string) Problem {
	return Problem{
		Type:     "about:blank",
		Title:    http.StatusText(statusCode),
		Status:   statusCode,
		Message:  message,
		Instance: instance,
		Guid:     guid,
	}
}

func SendProblemDetailJson(ctx *gin.Context, statusCode int, message, instance, guid string) {
	problem := generateProblemJson(statusCode, message, instance, guid)

	ctx.Header("Content-Type", "application/problem+json")
	ctx.JSON(statusCode, problem)
}

func SendProblemDetailJsonHttp(w http.ResponseWriter, statusCode int, message, instance, guid string) {
	problem := generateProblemJson(statusCode, message, instance, guid)

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(problem)
}

func SendProblemDetailJsonValidate(ctx *gin.Context, statusCode int, message, instance, guid string, errors validator.ValidationErrors) {
	errorKv := map[string]string{}

	for _, v := range errors {
		ns := v.Namespace()
		keys := strings.Split(ns, ".")

		errorKv[keys[1]] = normalizeError(v)
	}

	problem := generateProblemJson(statusCode, message, instance, guid)
	problem.Errors = errorKv

	ctx.Header("Content-Type", "application/problem+json")
	ctx.JSON(statusCode, problem)
}

func normalizeError(err validator.FieldError) string {
	//? every time using new validator, need to define it's error message here
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%v is required", strings.Split(err.Namespace(), ".")[1])
	case "max":
		return fmt.Sprintf("%v max length is %v", strings.Split(err.Namespace(), ".")[1], err.Param())
	case "gt":
		return fmt.Sprintf("%v value must be greater than %v", strings.Split(err.Namespace(), ".")[1], err.Param())
	case "gte":
		return fmt.Sprintf("%v value must be greater or equal than %v", strings.Split(err.Namespace(), ".")[1], err.Param())
	case "lte":
		return fmt.Sprintf("%v value must be lower or equal than %v", strings.Split(err.Namespace(), ".")[1], err.Param())
	default:
		return "undefined error"
	}
}
