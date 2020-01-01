package config

import (
	"gin-proyek/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=priantikonuradipratama password=admin dbname=gorm sslmode=disable")
	if err != nil {
		panic("database cannot connect, sorry")
	}

	DB.AutoMigrate(&models.Article{})
}
