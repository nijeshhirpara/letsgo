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
		response.Message = "Can't retrive users as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// Retrive users for the given company
	response.Data = b.UHandler.userRepo.ListUsersByCompany(company.ID)
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
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	// check if user's name is provided
	if uReq.Name == "" {
		response.Status = "error"
		response.Message = "User's name must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// check if user's email is provided
	if uReq.Email == "" {
		response.Status = "error"
		response.Message = "User's email must be supplied"
		SendJsonResponse(w, response)
		return
	}

	if user, err := b.UHandler.userRepo.FindUserByEmail(uReq.Email); err == nil && user.ID > 0 {
		response.Status = "error"
		response.Message = "User already exists with a given email"
		SendJsonResponse(w, response)
		return
	}

	user := models.User{
		Name:  uReq.Name,
		Email: uReq.Email,
	}

	if err := b.UHandler.userRepo.CreateUser(company, user); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
