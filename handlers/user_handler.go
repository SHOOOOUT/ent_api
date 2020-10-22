package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world")
}

func CreateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World POST")
}
