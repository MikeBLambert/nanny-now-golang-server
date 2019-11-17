package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mikeblambert/nanny-now-golang-server/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "nanny1",
		Email:    "nanny1@gmail.com",
		Password: "password1",
		Role:     "nanny",
	},
	models.User{
		Nickname: "family1",
		Email:    "family1@gmail.com",
		Password: "password2",
		Role:     "family",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
