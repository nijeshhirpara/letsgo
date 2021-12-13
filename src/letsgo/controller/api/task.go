package api

import (
	"encoding/json"
	"letsgo/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TaskHandler will hold everything that controller needs
type TaskHandler struct {
	taskRepo models.TaskRepository
}

// NewTaskHandler returns a new TaskHandler
func NewTaskHandler(taskRepo models.TaskRepository) *TaskHandler {
	return &TaskHandler{
		taskRepo: taskRepo,
	}
}

// CreateTask creates a task
func (b *BaseHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message`
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
		response.Message = "Can't add task as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	// Retrive team id parameter from url and validate
	teamID, err := strconv.ParseUint(vars["teamid"], 10, 64)
	if err != nil {
		response.Message = "Team ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// Find team by ID
	team, err := b.THandler.teamRepo.FindTeamByID(uint(teamID))
	if err != nil {
		response.Message = "Can't add task as given team doesn't exist"
		SendJsonResponse(w, response)
		return
	}
	if team.CompanyID != company.ID {
		response.Message = "Can't add task as given team doesn't exist in the company"
		SendJsonResponse(w, response)
		return
	}

	// request struct
	type TaskReq struct {
		Name   string `json:"name"`
		Status string `json:"status"`
		UserID int    `json:"user_id"`
	}
	var taskReq TaskReq

	// parse request body
	err = json.NewDecoder(r.Body).Decode(&taskReq)
	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	if taskReq.Name == "" {
		response.Status = "error"
		response.Message = "Task name must be supplied"
		SendJsonResponse(w, response)
		return
	}

	if taskReq.UserID == 0 {
		response.Status = "error"
		response.Message = "User ID must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// Find user by ID
	user, err := b.UHandler.userRepo.FindUserByID(uint(taskReq.UserID))
	if err != nil {
		response.Message = "Can't add task as given user doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	task := models.Task{
		Name:   taskReq.Name,
		Status: taskReq.Status,
	}

	if err := b.TskHandler.taskRepo.CreateTask(team, user, task); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}

// ListTasksByTeam retrives tasks for the given team
func (b *BaseHandler) ListTasksByTeam(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data    []models.Task `json:"data"`
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
		response.Message = "Can't retrive tasks as given company doesn't exist"
		SendJsonResponse(w, response)
		return
	}

	teamID, err := strconv.ParseUint(vars["teamid"], 10, 64)
	if err != nil {
		response.Message = "Team ID is not valid, it must be supplied"
		SendJsonResponse(w, response)
		return
	}

	// Find team by ID
	team, err := b.THandler.teamRepo.FindTeamByID(uint(teamID))
	if err != nil {
		response.Message = "Can't retrive tasks as given team doesn't exist"
		SendJsonResponse(w, response)
		return
	}
	if team.CompanyID != company.ID {
		response.Message = "Can't add user to teams as given team doesn't exist in the company"
		SendJsonResponse(w, response)
		return
	}

	// Retrieve tasks for the given team
	response.Data = b.TskHandler.taskRepo.ListTasksByTeam(team.ID)
	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
