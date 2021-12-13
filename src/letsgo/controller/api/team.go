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
	response.Data = b.THandler.teamRepo.ListTeamsByCompany(company.ID)
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

// AddUsersToTeam adds a list users to the team to within a company
func (b *BaseHandler) AddUsersToTeam(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message`
	}

	var response Response
	response.Status = "error"

	// retrive company id parameter from url and validate
	vars := mux.Vars(r)
	companyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Message = "Company ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if company exists with given ID
	company, err := b.CHandler.companyRepo.FindCompanyByID(uint(companyID))
	if err != nil {
		response.Message = "Can't add users to the team as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	teamID, err := strconv.ParseUint(vars["teamid"], 10, 64)
	if err != nil {
		response.Message = "Team ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if team exists with given ID
	team, err := b.THandler.teamRepo.FindTeamByID(uint(teamID))
	if err != nil {
		response.Message = "Can't add users to the team as given team doesn't exist"
		SendJsonResponse(w, response)
		return
	}
	if team.CompanyID != company.ID {
		response.Message = "Can't add users to the team as given team doesn't exist in the company"
		SendJsonResponse(w, response)
		return
	}

	// request struct
	type AddUsersReq struct {
		UsersID []int `json:"users_id"`
	}
	var addUsersReq AddUsersReq

	// parse request body
	err = json.NewDecoder(r.Body).Decode(&addUsersReq)
	if err != nil {
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	if len(addUsersReq.UsersID) == 0 {
		response.Message = "No users supplied"
		SendJsonResponse(w, response)
		return
	}

	count := len(addUsersReq.UsersID)

	for _, userID := range addUsersReq.UsersID {
		user, err := b.UHandler.userRepo.FindUserByID(uint(userID))
		if err != nil || user.CompanyID != company.ID {
			response.Message = "Incorrect user supplied"
			SendJsonResponse(w, response)
			return
		}

		teams := []models.Team{team}
		if err := b.UHandler.userRepo.AddTeamsToUser(user, teams); err != nil {
			count--
		}
	}

	if count <= 0 {
		response.Message = "Couldn't add any user to the team. Something went wrong"
		SendJsonResponse(w, response)
		return
	}

	if count < len(addUsersReq.UsersID) {
		response.Message = "Couldn't add few user to the team."
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
