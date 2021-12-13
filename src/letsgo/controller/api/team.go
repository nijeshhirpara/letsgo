package api

import (
	"encoding/json"
	"letsgo/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TeamHandler will hold everything that controller needs
type TeamHandler struct {
	teamRepo models.TeamRepository
}

// NewTeamHandler returns a new TeamHandler
func NewTeamHandler(teamRepo models.TeamRepository) *TeamHandler {
	return &TeamHandler{
		teamRepo: teamRepo,
	}
}

// ListTeams retrives teams for the given company
func (b *BaseHandler) ListTeams(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data    []models.Team `json:"data"`
		Status  string        `json:"status"`
		Message string        `json:"message"`
	}

	var response Response
	response.Status = "success"

	// Retrive company id parameter from url and validate
	vars := mux.Vars(r)
	companyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Status = "error"
		response.Message = "Company ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// Find company by ID
	company, err := b.CHandler.companyRepo.FindCompanyByID(uint(companyID))
	if err != nil {
		response.Status = "error"
		response.Message = "Can't retrive teams as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// Retrive teams for the given company
	response.Data = b.THandler.teamRepo.ListTeamsByCompany(uint(companyID))
	SendJsonResponse(w, response)
	return
}

// CreateTeam creates a team within a company if not exist
func (b *BaseHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message`
	}

	var response Response

	// retrive company id parameter from url and validate
	vars := mux.Vars(r)
	companyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Status = "error"
		response.Message = "Company ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if company exists with given ID
	company, err := b.CHandler.companyRepo.FindCompanyByID(uint(companyID))
	if err != nil {
		response.Status = "error"
		response.Message = "Can't retrive teams as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// request struct
	type TeamReq struct {
		Name        string `json:"name"`
		Description string `json:description`
	}
	var tReq TeamReq

	// parse request body
	err = json.NewDecoder(r.Body).Decode(&tReq)
	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	// check if company name is provided
	if tReq.Name == "" {
		response.Status = "error"
		response.Message = "Company name must be supplied"
		SendJsonResponse(w, response)
		return
	}

	if team, err := b.THandler.teamRepo.FindCompanyTeamByName(company.ID, tReq.Name); err == nil && team.ID > 0 {
		response.Status = "error"
		response.Message = "Team already exists with a given name"
		SendJsonResponse(w, response)
		return
	}

	team := models.Team{
		Name:        tReq.Name,
		Description: tReq.Description,
	}

	if err := b.CHandler.companyRepo.AddTeamToCompany(company, team); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
