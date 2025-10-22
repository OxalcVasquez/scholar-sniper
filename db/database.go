package db

import (
	"log"

	"github.com/oxalcvasquez/scholar-sni/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("scholarships.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = DB.AutoMigrate(&model.Scholarchip{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

}
