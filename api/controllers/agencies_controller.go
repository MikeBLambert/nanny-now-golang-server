package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mikeblambert/nanny-now-golang-server/api/auth"
	"github.com/mikeblambert/nanny-now-golang-server/api/models"
	"github.com/mikeblambert/nanny-now-golang-server/api/responses"
	"github.com/mikeblambert/nanny-now-golang-server/api/utils/formaterror"
)

func (server *Server) CreateAgency(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	agency := models.Agency{}
	err = json.Unmarshal(body, &agency)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	agency.Prepare()
	err = agency.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = auth.ExtractTokenID(request)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	agencyCreated, err := agency.SaveAgency(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", request.Host, request.URL.Path, agencyCreated.ID))
	responses.JSON(w, http.StatusCreated, agencyCreated)
}

func (server *Server) GetAgencies(w http.ResponseWriter, request *http.Request) {

	agency := models.Agency{}

	agencies, err := agency.FindAllAgencies(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, agencies)
}

func (server *Server) GetAgency(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	agencyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	agency := models.Agency{}

	agencyReceived, err := agency.FindAgencyByID(server.DB, agencyID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, agencyReceived)
}

func (server *Server) UpdateAgency(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	// Check if the agency id is valid
	agencyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	_, err = auth.ExtractTokenID(request)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the agency exist
	agency := models.Agency{}
	err = server.DB.Debug().Model(models.Agency{}).Where("id = ?", agencyID).Take(&agency).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Agency not found"))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	agencyUpdate := models.Agency{}
	err = json.Unmarshal(body, &agencyUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	agencyUpdate.Prepare()
	err = agencyUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	agencyUpdate.ID = agency.ID //this is important to tell the model the agency id to update, the other update field are set above

	agencyUpdated, err := agencyUpdate.UpdateAnAgency(server.DB, agencyID)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, agencyUpdated)
}

func (server *Server) DeleteAgency(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	agencyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	_, err = auth.ExtractTokenID(request)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the agency exists
	agency := models.Agency{}
	err = server.DB.Debug().Model(models.Agency{}).Where("id = ?", agencyID).Take(&agency).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = agency.DeleteAnAgency(server.DB, agencyID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", agencyID))
	responses.JSON(w, http.StatusNoContent, "")
}
