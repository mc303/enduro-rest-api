package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Create datbase models using ORM
func Create(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Post{})
	// db.DropTableIfExists(&Rider{}, &Class{}, &Run{}, &Stage{}, &Event{}, &TypeOfRace{})
	// set up foreign keys
	// db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	c := []Class{
		{Name: "Masster"},
		{Name: "Junior"},
		{Name: "Senior"},
	}

	for i := range c {
		if err := db.Debug().Save(&c[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	tor := []TypeOfRace{
		{Name: "Cross Country"},
		{Name: "Cross Marathon"},
		{Name: "Downhill"},
		{Name: "Enduro"},
		{Name: "Four Cross"},
	}

	for i := range tor {
		if err := db.Save(&tor[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	ev := []Event{
		{
			Name:        "Enduro 01",
			TotalStages: 3,
			TypeOfRaces: tor[3],
		},
		{
			Name:        "Enduro 02",
			TotalStages: 3,
			TypeOfRaces: tor[3],
		},
	}
	for i := range ev {
		if err := db.Save(&ev[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// type Stage struct {
	// 	ID       int
	//	Name 	 string	`gorm:"not null"`
	// 	Events   Event
	// 	EventID  uint `gorm:"foreignkey:id"`
	// 	Location string
	// 	Order     bool
	// }
	s := []Stage{
		{
			Name:   "Stage 0001",
			Order:  1,
			Events: ev[0],
		},
		{
			Name:   "Stage 0002",
			Order:  2,
			Events: ev[0],
		},
		{
			Name:   "Stage 0003",
			Order:  3,
			Events: ev[0],
		},
		{
			Name:   "BB Stage 0001",
			Order:  1,
			Events: ev[1],
		},
		{
			Name:   "BB Stage 0002",
			Order:  2,
			Events: ev[1],
		},
		{
			Name:   "BB Stage 0003",
			Order:  3,
			Events: ev[1],
		},
	}

	for i := range s {
		if err := db.Save(&s[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// Db.Save(&c) //Saving one to one relationship
	// Db.Save(&c2)
	// Db.Save(&c3)

	rider := []Rider{
		{
			Firstname: "John",
			Lastname:  "Doe",
			Gender:    2,
			Mail:      "john.doe@gmail.com",
			Class:     c[0],
		},
		{
			Firstname: "Peter",
			Lastname:  "Peterersen",
			Gender:    2,
			Mail:      "peter.peterersen@gmail.com",
			Class:     c[0],
		},
		{
			Firstname: "Jane",
			Lastname:  "Doe",
			Gender:    2,
			Mail:      "jane.doe@gmail.com",
			Class:     c[1],
		},
	}

	for i := range rider {
		if err := db.Save(&rider[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	// Rider makes a run
	r := []Run{
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[0],
			Stages:  s[0],
		},
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[1],
			Stages:  s[0],
		},
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[1],
			Stages:  s[1],
		},
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[0],
			Stages:  s[1],
		},
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[1],
			Stages:  s[2],
		},
		{StartTime: time.Now(),
			EndTime: time.Now(),
			DNF:     false,
			Riders:  rider[0],
			Stages:  s[2],
		},
	}

	for i := range r {
		if err := db.Save(&r[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	se := []Season{
		{
			Name: "Season 2018",
			Year: 2018,
		},
		{
			Name: "Season 2019",
			Year: 2019,
		},
	}

	for i := range se {
		if err := db.Save(&se[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Created database has beed processed")
}
