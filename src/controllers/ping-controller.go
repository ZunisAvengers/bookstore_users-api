package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.String(http.StatusOK, "Pong")
}

func GetFavicon(context *gin.Context) {
	body, err := os.ReadFile("assets/favicon.ico")
	if err != nil {
		context.Err()

	}
	context.Writer.Write(body)
}
