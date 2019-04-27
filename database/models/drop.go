package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Drop datbase models using ORM
func Drop(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Post{})
	db.DropTableIfExists(&Rider{}, &Class{}, &Run{}, &Stage{}, &Event{}, &TypeOfRace{}, &Season{}, &Registered{}, &Result{}, &User{}, &UserAccess{})
	// set up foreign keys
	// db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Drop database has beed processed")
}
