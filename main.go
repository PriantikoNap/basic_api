package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text;"`
}

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=priantikonuradipratama password=admin dbname=gorm sslmode=disable")
	if err != nil {
		panic("database cannot connect, sorry")
	}
	defer DB.Close()
	DB.AutoMigrate(&Article{})

	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/about", getAbout)
	router.GET("/content/:title", getContent)
	router.POST("/contents", postContent)

	router.Run()
}

func getHome(c *gin.Context) {
	items := []Article{}
	DB.Find(&items)
	c.JSON(200, gin.H{
		"message": "Done in Home",
		"data":    items,
	})
}
func getAbout(c *gin.Context) {
	//tiko : this is about feature, thanks
	c.JSON(200, gin.H{
		"message": "Done this is About",
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

func postContent(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	c.JSON(200, gin.H{
		"status": "berhasil",
		"title":  title,
		"desc":   desc,
	})
}
