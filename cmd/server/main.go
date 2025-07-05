package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/ping", serverCheck)
	router.Run(":8000")

}

func serverCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "pong")
}
