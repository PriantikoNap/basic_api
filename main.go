package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", getHome)

	router.Run()
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Berhasil di Home",
	})
}
