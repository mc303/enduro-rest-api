package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// CreateClass add items to table
// func CreateClass(db *gorm.DB) *[]Class {

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
	starttime := time.Now()

	c := []Category{
		{Name: "Master",
			Gender: 1,
		},
		{Name: "Junior",
			Gender: 1,
		},
		{Name: "Senior",
			Gender: 1,
		},
		{Name: "Master E-MTB",
			Gender: 1,
		},
		{Name: "Junior E-MTB",
			Gender: 1,
		},
		{Name: "Senior E-MTB",
			Gender: 1,
		},
	}

	for i := range c {
		if err := db.Debug().Save(&c[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	tor := []Discipline{
		{Name: "ross Country"},
		{Name: "Cross Marathon"},
		{Name: "Downhill"},
		{Name: "Enduro"},
		{Name: "4X Cross"},
	}

	for i := range tor {
		if err := db.Save(&tor[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	ev := []Event{
		{
			Name: "BIKE Festival Willingen 2019",
			Date: starttime,
		},
		{
			Name: "LichtVerzet Festival 2019",
			Date: starttime,
		},
	}
	for i := range ev {
		if err := db.Save(&ev[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// type Stage struct {
	// 	ID       int
	// 	Name 	 string	`gorm:"not null"`
	// 	Events   Event
	// 	EventID  uint `gorm:"foreignkey:id"`
	// 	Location string
	// 	Order     bool
	// )

	s := []Stage{
		{
			Name:    "Stage 0001",
			Order:   1,
			EventID: 1,
			RaceID:  1,
		},
		{
			Name:    "Stage 0002",
			Order:   2,
			EventID: 1,
			RaceID:  1,
		},
		{
			Name:    "Stage 0003",
			Order:   3,
			EventID: 1,
			RaceID:  1,
		},
		{
			Name:    "BB Stage 0001",
			Order:   1,
			EventID: 1,
			RaceID:  2,
		},
		{
			Name:    "BB Stage 0002",
			Order:   2,
			EventID: 1,
			RaceID:  2,
		},
		{
			Name:    "BB Stage 0003",
			Order:   3,
			EventID: 1,
			RaceID:  2,
		},
	}

	for i := range s {
		if err := db.Save(&s[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	address := []Address{
		{
			Street:       "De Weg",
			StreetNumber: "1",
			PostalCode:   "7777MM",
			City:         "Venlo",
			Country:      "Nederland",
			EventID:      1,
		},
		{
			Street:       "De Straat",
			StreetNumber: "66",
			PostalCode:   "2487 GA",
			City:         "Utrecht",
			Country:      "NL",
			RiderID:      1,
		},
		{
			Street:       "De Straatweg",
			StreetNumber: "65",
			PostalCode:   "3456 FR",
			City:         "Rotterdam",
			Country:      "Netherlands",
			RiderID:      2,
		},
	}

	for i := range address {
		if err := db.Save(&address[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// Db.Save(&c) //Saving one to one relationship
	// Db.Save(&c2)
	// Db.Save(&c3)

	ri := []Rider{
		{
			Firstname: "John",
			Lastname:  "Doe",
			Gender:    1,
			Mail:      "john.doe@gmail.com",
			
			// Address:	address[1],
		},
		{
			Firstname: "Peter",
			Lastname:  "Petersen",
			Gender:    1,
			Mail:      "peter.petersen@gmail.com",
		},
		{
			Firstname: "Jane",
			Lastname:  "Doe",
			Gender:    1,
			Mail:      "jane.doe@gmail.com",
		},
	}

	for i := range ri {
		if err := db.Save(&ri[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	//Rider makes a run

	r := []Run{
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 32),
			DNF:     false,
			RiderID: 1,
			StageID: 1,
			EventID: 1,
		},
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 33),
			DNF:     false,
			RiderID: 2,
			StageID: 1,
			EventID: 1,
		},
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 34),
			DNF:     false,
			RiderID: 1,
			StageID: 2,
			EventID: 1,
		},
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 32),
			DNF:     false,
			RiderID: 2,
			StageID: 2,
			EventID: 1,
		},
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 33),
			DNF:     false,
			RiderID: 1,
			StageID: 3,
			EventID: 1,
		},
		{StartTime: starttime,
			EndTime: starttime.Add(time.Second * 35),
			DNF:     false,
			RiderID: 2,
			StageID: 3,
			EventID: 1,
		},
	}

	for i := range r {
		if err := db.Save(&r[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// "ID":        p.ID,
	// "TotalTime": p.TotalTime,
	// "Riders":    p.Riders,
	// "RidersID":  p.RidersID,
	// "Events":    p.Events,
	// "EventsID":  p.EventsID,
	// "Seasons":   p.Seasons,
	// "SeasonsID": p.SeasonsID,
	// "Place":     p.Place,

	result := []Result{
		{
			TotalTime: time.Now(),
			RiderID:   1,
			EventID:   1,
			Place:     1,
		},
		{
			TotalTime: time.Now(),
			RiderID:   2,
			EventID:   1,
			Place:     2,
		},
		// {
		// 	TotalTime: time.Now(),
		// 	RiderID:    0,
		// 	EventID:    0,
		// 	Place:     3,
		// },
	}

	for i := range result {
		if err := db.Save(&result[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Created database has beed processed")
}
