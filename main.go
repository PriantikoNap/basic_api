package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/about", getAbout)
	router.GET("/content/:title", getContent)

	router.Run()
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Done in Home",
	})
}
func getAbout(c *gin.Context) {
	//tiko : this is about feature, thanks
	c.JSON(200, gin.H{
		"message": "Done in About",
		"no-telp": "get telp number",
	})
}
func getContent(c *gin.Context) {
	title := c.Param("title")
	//tiko: this is get from db api, thanks
	c.JSON(200, gin.H{
		"message": title,
		"status":  "Ok",
	})
}
