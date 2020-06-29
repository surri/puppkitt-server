package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"puppkitt.com/v1/database/models"
	_"github.com/go-sql-driver/mysql"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")

	fmt.Println(dbConfig)
	db, err := gorm.Open("mysql", dbConfig)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
