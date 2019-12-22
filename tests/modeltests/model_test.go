package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/mikeblambert/nanny-now-golang-server/api/controllers"
	"github.com/mikeblambert/nanny-now-golang-server/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}
var agencyInstance = models.Agency{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshAgencyTable() error {
	err := server.DB.DropTableIfExists(&models.Agency{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate((&models.Agency{})).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed agency table")
	return nil
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user table")
	return nil
}

func refreshAgencyAndUserTable() error {
	err := server.DB.DropTableIfExists(&models.Agency{}, &models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.Agency{}, &models.User{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneAgency() (models.Agency, error) {

	refreshAgencyTable()

	agency := models.Agency{
		BusinessName:   "Nannies 'R Us",
		ContactName:    "Mike Mikerson",
		StreetAddress1: "123 Main St",
		StreetAddress2: "Apt 3",
		State:          "NY",
	}

	err := server.DB.Model(&models.Agency{}).Create(&agency).Error
	if err != nil {
		log.Fatalf("cannot seed agencies table: %v", err)
	}
	return agency, nil
}

func seedAgencies() error {

	refreshAgencyTable()

	agencies := []models.Agency{
		models.Agency{
			BusinessName:   "Nannies 'R Us",
			ContactName:    "Mike Mikerson",
			StreetAddress1: "123 Main St",
			StreetAddress2: "Apt 3",
			State:          "NY",
		},
		models.Agency{
			BusinessName:   "Big 'ol Babies",
			ContactName:    "Baby McBabyson",
			StreetAddress1: "123 Front St",
			StreetAddress2: "Apt 8",
			State:          "OR",
		},
	}

	for i := range agencies {
		err := server.DB.Model(&models.Agency{}).Create(&agencies[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedOneUserAndOneAgency() (models.User, error) {
	err := refreshAgencyAndUserTable()
	if err != nil {
		return models.User{}, err
	}

	agency := models.Agency{
		BusinessName:   "Nannies 'R Us",
		ContactName:    "Mike Mikerson",
		StreetAddress1: "123 Main St",
		StreetAddress2: "Apt 3",
		State:          "NY",
	}

	err = server.DB.Model(&models.Agency{}).Create(&agency).Error
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Nickname: "Jane",
		Email:    "jane@gmail.com",
		Password: "password",
		Role:     "nanny",
		Agency:   agency.ID,
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Nickname: "Steven",
			Email:    "steven@gmail.com",
			Password: "password",
			Role:     "nanny",
			Agency:   1,
		},
		models.User{
			Nickname: "Kenny",
			Email:    "kenny@gmail.com",
			Password: "password",
			Role:     "family",
			Agency:   1,
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
