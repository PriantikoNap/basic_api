package main

import (
	"gin-proyek/config"
	"gin-proyek/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()
	router.GET("/", routes.GetHome)
	router.GET("/about", routes.GetAbout)
	router.GET("/content/:slug", routes.GetContent)
	router.POST("/contents", routes.PostContent)

	router.Run()
}
