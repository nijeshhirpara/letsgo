package api

import (
	"encoding/json"
	"letsgo/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler will hold everything that controller needs
type UserHandler struct {
	userRepo models.UserRepository
}

// NewUserHandler returns a new UserHandler
func NewUserHandler(userRepo models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// ListTeams retrives users for the given company
func (b *BaseHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data    []models.User `json:"data"`
		Status  string        `json:"status"`
		Message string        `json:"message"`
	}

	var response Response
	response.Status = "error"

	// Retrive company id parameter from url and validate
	vars := mux.Vars(r)
	companyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.Message = "Company ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// Find company by ID
	company, err := b.CHandler.companyRepo.FindCompanyByID(uint(companyID))
	if err != nil {
		response.Message = "Can't retrive users as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// Retrive users for the given company
	response.Data = b.UHandler.userRepo.ListUsersByCompany(company.ID)
	response.Status = "success"
	SendJsonResponse(w, response)
	return
}

// CreateTeam creates a user within a company if not exist
func (b *BaseHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
		response.Message = "Can't retrive users as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// request struct
	type UserReq struct {
		Name  string `json:"name"`
		Email string `json:email`
	}
	var uReq UserReq

	// parse request body
	err = json.NewDecoder(r.Body).Decode(&uReq)
	if err != nil {
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	// check if user's name is provided
	if uReq.Name == "" {
		response.Message = "User's name must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if user's email is provided
	if uReq.Email == "" {
		response.Message = "User's email must be supplied"
		SendJsonResponse(w, response)
		return
	}

	if user, err := b.UHandler.userRepo.FindUserByEmail(uReq.Email); err == nil && user.ID > 0 {
		response.Message = "User already exists with a given email"
		SendJsonResponse(w, response)
		return
	}

	user := models.User{
		Name:  uReq.Name,
		Email: uReq.Email,
	}

	if err := b.UHandler.userRepo.CreateUser(company, user); err != nil {
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}

// AddUserToTeams adds a user to list of teams to within a company
func (b *BaseHandler) AddUserToTeams(w http.ResponseWriter, r *http.Request) {
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
		response.Message = "Can't add user to teams as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	userID, err := strconv.ParseUint(vars["userid"], 10, 64)
	if err != nil {
		response.Message = "User ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if user exists with given ID
	user, err := b.UHandler.userRepo.FindUserByID(uint(userID))
	if err != nil {
		response.Message = "Can't add user to teams as given user doesn't exist"
		SendJsonResponse(w, response)
		return
	}
	if user.CompanyID != company.ID {
		response.Message = "Can't add user to teams as given user doesn't exist in the company"
		SendJsonResponse(w, response)
		return
	}

	// request struct
	type AddUserReq struct {
		TeamsID []int `json:"teams_id"`
	}
	var addUserReq AddUserReq

	// parse request body
	err = json.NewDecoder(r.Body).Decode(&addUserReq)
	if err != nil {
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	if len(addUserReq.TeamsID) == 0 {
		response.Message = "No teams supplied"
		SendJsonResponse(w, response)
		return
	}

	var teams []models.Team

	for _, teamID := range addUserReq.TeamsID {
		team, err := b.THandler.teamRepo.FindTeamByID(uint(teamID))
		if err != nil || team.CompanyID != company.ID {
			response.Message = "Incorrect team supplied"
			SendJsonResponse(w, response)
			return
		}
		teams = append(teams, team)
	}

	if err := b.UHandler.userRepo.AddTeamsToUser(user, teams); err != nil {
		response.Message = "Couldn't add user to teams. Something went wrong"
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
