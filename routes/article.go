package routes

import (
	"gin-proyek/config"
	"gin-proyek/models"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	config.DB.Find(&items)
	c.JSON(200, gin.H{
		"message": "Done in Home",
		"data":    items,
	})
}
func GetAbout(c *gin.Context) {
	//tiko : this is about feature, thanks
	c.JSON(200, gin.H{
		"message": "Done this is About",
		"no-telp": "get telp number",
	})
}
func GetContent(c *gin.Context) {
	slug := c.Param("slug")
	var item models.Article
	if config.DB.First(&item, "slug=?", slug).RecordNotFound() {
		c.JSON(404, gin.H{
			"status": "error",
			"msg":    "record not found",
		})
		c.Abort()
		return
	}

	//tiko: this is get from db api, thanks
	c.JSON(200, gin.H{
		"data":   item,
		"status": "Ok",
	})
}

func PostContent(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  c.PostForm("slug"),
	}
	config.DB.Create(&item)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}
