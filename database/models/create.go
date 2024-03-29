package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Create datbase models using ORM
func Create(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Post{})
	// db.DropTableIfExists(&Rider{}, &Class{}, &Run{}, &Stage{}, &Event{}, &TypeOfRace{})
	// set up foreign keys
	// db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// db.Model(&Address{}).AddForeignKey("riders_id", "riders(id)", "CASCADE", "CASCADE")
	// db.Model(&Rider{}).Related(&Address{})
	// var adres models.Address
	// db.Model(&Rider).Related(&adres, "Address")

	// c := []Class{
	// 	{Name: "Masster"},
	// 	{Name: "Junior"},
	// 	{Name: "Senior"},
	// }

	// for i := range c {
	// 	if err := db.Debug().Save(&c[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// tor := []TypeOfRace{
	// 	{Name: "Cross Country"},
	// 	{Name: "Cross Marathon"},
	// 	{Name: "Downhill"},
	// 	{Name: "Downhill E-MTB"},
	// 	{Name: "Enduro"},
	// 	{Name: "Enduro E-MTB"},
	// 	{Name: "Four Cross"},
	// }

	// se := []Season{
	// 	{
	// 		Name: "Season 2018",
	// 		Year: 2018,
	// 	},
	// 	{
	// 		Name: "Season 2019",
	// 		Year: 2019,
	// 	},
	// }

	// for i := range se {
	// 	if err := db.Save(&se[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// for i := range tor {
	// 	if err := db.Save(&tor[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// ev := []Event{
	// 	{
	// 		Name:        "Enduro 01",
	// 		TotalStages: 3,
	// 		TypeOfRaces: tor[3],
	// 		Seasons:     se[0],
	// 	},
	// 	{
	// 		Name:        "Enduro 02",
	// 		TotalStages: 3,
	// 		TypeOfRaces: tor[3],
	// 		Seasons:     se[0],
	// 	},
	// }
	// for i := range ev {
	// 	if err := db.Save(&ev[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // type Stage struct {
	// // 	ID       int
	// //	Name 	 string	`gorm:"not null"`
	// // 	Events   Event
	// // 	EventID  uint `gorm:"foreignkey:id"`
	// // 	Location string
	// // 	Order     bool
	// // }
	// // s := []Stage{
	// // 	{
	// // 		Name:   "Stage 0001",
	// // 		Order:  1,
	// // 		Events: ev[0],
	// // 	},
	// // 	{
	// // 		Name:   "Stage 0002",
	// // 		Order:  2,
	// // 		Events: ev[0],
	// // 	},
	// // 	{
	// // 		Name:   "Stage 0003",
	// // 		Order:  3,
	// // 		Events: ev[0],
	// // 	},
	// // 	{
	// // 		Name:   "BB Stage 0001",
	// // 		Order:  1,
	// // 		Events: ev[1],
	// // 	},
	// // 	{
	// // 		Name:   "BB Stage 0002",
	// // 		Order:  2,
	// // 		Events: ev[1],
	// // 	},
	// // 	{
	// // 		Name:   "BB Stage 0003",
	// // 		Order:  3,
	// // 		Events: ev[1],
	// // 	},
	// // }

	// // for i := range s {
	// // 	if err := db.Save(&s[i]).Error; err != nil {
	// // 		log.Fatal(err)
	// // 	}
	// // }

	// // Db.Save(&c) //Saving one to one relationship
	// // Db.Save(&c2)
	// // Db.Save(&c3)

	// ri := []Rider{
	// 	{
	// 		Firstname: "John",
	// 		Lastname:  "Doe",
	// 		Gender:    2,
	// 		Mail:      "john.doe@gmail.com",
	// 		Class:     c[0],
	// 	},
	// 	{
	// 		Firstname: "Peter",
	// 		Lastname:  "Petersen",
	// 		Gender:    2,
	// 		Mail:      "peter.petersen@gmail.com",
	// 		Class:     c[0],
	// 	},
	// 	{
	// 		Firstname: "Jane",
	// 		Lastname:  "Doe",
	// 		Gender:    2,
	// 		Mail:      "jane.doe@gmail.com",
	// 		Class:     c[1],
	// 	},
	// }

	// for i := range ri {
	// 	if err := db.Save(&ri[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // Rider makes a run
	// // r := []Run{
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[0],
	// // 		Stages:  s[0],
	// // 	},
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[1],
	// // 		Stages:  s[0],
	// // 	},
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[1],
	// // 		Stages:  s[1],
	// // 	},
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[0],
	// // 		Stages:  s[1],
	// // 	},
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[1],
	// // 		Stages:  s[2],
	// // 	},
	// // 	{StartTime: time.Now(),
	// // 		EndTime: time.Now(),
	// // 		DNF:     false,
	// // 		Riders:  ri[0],
	// // 		Stages:  s[2],
	// // 	},
	// // }

	// for i := range r {
	// 	if err := db.Save(&r[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// address := []Address{
	// 	{
	// 		Street:       "De Weg",
	// 		StreetNumber: "1",
	// 		PostalCode:   "7777MM",
	// 		City:         "Venlo",
	// 		Country:      "Nederland",
	// 		Riders:       ri[0],
	// 	},
	// 	{
	// 		Street:       "De Straat",
	// 		StreetNumber: "66",
	// 		PostalCode:   "2487 GA",
	// 		City:         "Utrecht",
	// 		Country:      "NL",
	// 		Riders:       ri[1],
	// 	},
	// 	{
	// 		Street:       "De Straatweg",
	// 		StreetNumber: "65",
	// 		PostalCode:   "3456 FR",
	// 		City:         "Rotterdam",
	// 		Country:      "Netherlands",
	// 		Riders:       ri[2],
	// 	},
	// }

	// for i := range address {
	// 	if err := db.Save(&address[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // "ID":        p.ID,
	// // "TotalTime": p.TotalTime,
	// // "Riders":    p.Riders,
	// // "RidersID":  p.RidersID,
	// // "Events":    p.Events,
	// // "EventsID":  p.EventsID,
	// // "Seasons":   p.Seasons,
	// // "SeasonsID": p.SeasonsID,
	// // "Place":     p.Place,

	// result := []Result{
	// 	{
	// 		TotalTime: time.Now(),
	// 		Riders:    ri[0],
	// 		Events:    ev[0],
	// 		Place:     1,
	// 	},
	// 	{
	// 		TotalTime: time.Now(),
	// 		Riders:    ri[0],
	// 		Events:    ev[0],
	// 		Place:     2,
	// 	},
	// 	{
	// 		TotalTime: time.Now(),
	// 		Riders:    ri[0],
	// 		Events:    ev[0],
	// 		Place:     3,
	// 	},
	// }

	// for i := range result {
	// 	if err := db.Save(&result[i]).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	fmt.Println("Created database has beed processed")
}
