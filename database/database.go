package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/mc303/gin-rest-api-sample/database/models"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	//db, err := gorm.Open("mysql", dbConfig)
	db, err := gorm.Open("sqlite3", dbConfig)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Drop(db)
	models.Migrate(db)
	// models.Create(db)
	return db, err
}
