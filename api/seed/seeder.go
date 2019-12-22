package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mikeblambert/nanny-now-golang-server/api/models"
)

var agencies = []models.Agency{
	models.Agency{
		BusinessName:   "NW Nannies Inc",
		ContactName:    "Linda Roffe",
		StreetAddress1: "1286 Hide-a-way Ln.",
		StreetAddress2: "",
		State:          "OR",
	},
}

var users = []models.User{
	models.User{
		Nickname: "nanny1",
		Email:    "nanny1@gmail.com",
		Password: "password1",
		Role:     "nanny",
		Agency:   1,
	},
	models.User{
		Nickname: "family1",
		Email:    "family1@gmail.com",
		Password: "password2",
		Role:     "family",
		Agency:   1,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Agency{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Agency{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range agencies {
		err = db.Debug().Model(&models.Agency{}).Create(&agencies[i]).Error
		if err != nil {
			log.Fatalf("cannot seed agencies table: %v", err)
		}
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
