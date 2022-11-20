package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	context.String(http.StatusBadRequest, "CreateUser method")
}

func GetUsers(context *gin.Context) {
	context.String(http.StatusBadRequest, "GetUsers method")
}

// func SearchUser(context *gin.Context) {
// 	context.String(http.StatusBadRequest, "SearchUser method")
// }
