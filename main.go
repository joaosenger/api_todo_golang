package main

import (
	"api/config"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	config.DB = config.InitDB("Tasks.db")
	defer config.DB.Close()

	router.SetTrustedProxies(nil)

	router.Run(":3000")
}
