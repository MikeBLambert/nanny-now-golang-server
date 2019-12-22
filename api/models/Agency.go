package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Agency struct {
	ID             uint64 `gorm:"primary_key; auto_increment" json:"id"`
	BusinessName   string `gorm:"size:255; not null;" json:"businessName"`
	ContactName    string `gorm:"size:255; not null;" json:"contactName"`
	StreetAddress1 string `gorm:"not null;" json:"streetAddress1"`
	StreetAddress2 string `gorm:"" json:"streetAddress2"`
	State          string `gorm:"size:255; not null" json:"state"`
}

func (a *Agency) Prepare() {
	a.ID = 0
	a.BusinessName = html.EscapeString(strings.TrimSpace(a.BusinessName))
	a.ContactName = html.EscapeString(strings.TrimSpace((a.ContactName)))
	a.StreetAddress1 = html.EscapeString(strings.TrimSpace((a.StreetAddress1)))
	a.StreetAddress2 = html.EscapeString(strings.TrimSpace((a.StreetAddress2)))
	a.State = html.EscapeString(strings.TrimSpace((a.State)))

}

func (a *Agency) Validate() error {
	if a.BusinessName == "" {
		return errors.New("Business Name Required")
	}
	if a.ContactName == "" {
		return errors.New("Contact Name Required")
	}
	if a.StreetAddress1 == "" {
		return errors.New("Street Address Line 1 Required")
	}
	if a.State == "" {
		return errors.New("State Required")
	}
	return nil
}

func (a *Agency) SaveAgency(db *gorm.DB) (*Agency, error) {
	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Agency{}, err
	}
	return a, nil
}

func (a *Agency) FindAllAgencies(db *gorm.DB) (*[]Agency, error) {
	var err error
	agencies := []Agency{}
	err = db.Debug().Model(&Agency{}).Limit(100).Find(&agencies).Error
	if err != nil {
		return &[]Agency{}, err
	}
	return &agencies, err
}

func (a *Agency) FindAgencyByID(db *gorm.DB, agencyID uint64) (*Agency, error) {
	var err error
	err = db.Debug().Model(Agency{}).Where("id = ?", agencyID).Take(&a).Error
	if err != nil {
		return &Agency{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Agency{}, errors.New("Agency Not Found")
	}
	return a, err
}

func (a *Agency) UpdateAnAgency(db *gorm.DB, agencyID uint64) (*Agency, error) {
	var err error
	err = db.Debug().Model(&Agency{}).Where("id = ?", a.ID).Updates(Agency{BusinessName: a.BusinessName, ContactName: a.ContactName, StreetAddress1: a.StreetAddress1, StreetAddress2: a.StreetAddress2, State: a.State}).Error

	if err != nil {
		return &Agency{}, err
	}
	return a, nil
}

func (a *Agency) DeleteAnAgency(db *gorm.DB, agencyID uint64) (int64, error) {
	db = db.Debug().Model(&Agency{}).Where("id = ?", agencyID).Take(&Agency{}).Delete(&Agency{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Agency not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
