package modeltests

import "testing"

import "log"

import "gopkg.in/go-playground/assert.v1"

import "github.com/mikeblambert/nanny-now-golang-server/api/models"

func TestFindAllAgencies(t *testing.T) {
	err := refreshAgencyTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedAgencies()
	if err != nil {
		log.Fatal(err)
	}

	agencies, err := agencyInstance.FindAllAgencies(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the agencies: %v\n", err)
		return
	}
	assert.Equal(t, len(*agencies), 2)
}

func TestSaveAgency(t *testing.T) {
	err := refreshAgencyTable()
	if err != nil {
		log.Fatal(err)
	}
	newAgency := models.Agency{
		BusinessName:   "Test Business Name",
		ContactName:    "Test Contact Name",
		StreetAddress1: "Test Street Address 1",
		StreetAddress2: "Test Street Address 2",
		State:          "Test State",
	}

	savedAgency, err := newAgency.SaveAgency(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the agencies: %v\n", err)
		return
	}
	assert.Equal(t, newAgency.ID, savedAgency.ID)
	assert.Equal(t, newAgency.BusinessName, savedAgency.BusinessName)
	assert.Equal(t, newAgency.ContactName, savedAgency.ContactName)
	assert.Equal(t, newAgency.StreetAddress1, savedAgency.StreetAddress1)
	assert.Equal(t, newAgency.StreetAddress2, savedAgency.StreetAddress2)
	assert.Equal(t, newAgency.State, savedAgency.State)
}

func TestGetAgencyByID(t *testing.T) {
	err := refreshAgencyTable()
	if err != nil {
		log.Fatal(err)
	}

	agency, err := seedOneAgency()
	if err != nil {
		log.Fatalf("cannot seed agencies table: %v", err)
	}

	foundAgency, err := agencyInstance.FindAgencyByID(server.DB, agency.ID)
	if err != nil {
		t.Errorf("this is the error getting one agency: %v\n", err)
		return
	}

	assert.Equal(t, foundAgency.ID, agency.ID)
	assert.Equal(t, foundAgency.BusinessName, agency.BusinessName)
	assert.Equal(t, foundAgency.ContactName, agency.ContactName)
	assert.Equal(t, foundAgency.StreetAddress1, agency.StreetAddress1)
	assert.Equal(t, foundAgency.StreetAddress2, agency.StreetAddress2)
	assert.Equal(t, foundAgency.State, agency.State)
}
