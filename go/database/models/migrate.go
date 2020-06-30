package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Place{})
	// set up foreign keys
	db.Model(&Post{}).AddForeignKey("UserId", "users(Id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
