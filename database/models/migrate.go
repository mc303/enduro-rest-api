package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Post{})

	db.AutoMigrate(&Rider{}, &Class{}, &Run{}, &Stage{}, &Event{}, &TypeOfRace{}, &Season{}, &Registered{}, &Result{}, &User{}, &UserAccess{})

	// set up foreign keys
	// db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}
