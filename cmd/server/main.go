package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mateeusferro/fineval/internal/config"
	"github.com/mateeusferro/fineval/internal/delivery"
)

func main() {
	config.LoadEnv()
	var router *gin.Engine = gin.Default()
	port := config.EnvVariable("PORT")

	delivery.Routes(router)
	err := router.Run(":" + port)

	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}

}
