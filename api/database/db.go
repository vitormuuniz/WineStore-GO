package database

import (
	"log"

	"github.com/vitormuuniz/winestore-go/api/database/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.BuildDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
