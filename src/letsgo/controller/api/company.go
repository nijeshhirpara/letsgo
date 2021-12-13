package api

import (
	"encoding/json"
	"letsgo/models"
	"net/http"
)

// CompanyHandler will hold everything that controller needs
type CompanyHandler struct {
	companyRepo models.CompanyRepository
}

// NewCompanyHandler returns a new CompanyHandler
func NewCompanyHandler(companyRepo models.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{
		companyRepo: companyRepo,
	}
}

// ListCompanies retrives all companies from database
func (b *BaseHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data   []models.Company `json:"data"`
		Status string           `json:"status"`
	}

	var response Response
	response.Status = "success"
	response.Data = b.CHandler.companyRepo.ListCompanies()

	SendJsonResponse(w, response)
	return
}

// CreateCompany creates a company (company name must be unique)
func (b *BaseHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message`
	}

	var response Response

	// request struct
	type CompanyReq struct {
		Name        string `json:"name"`
		Description string `json:description`
	}
	var cReq CompanyReq

	// parse request body
	err := json.NewDecoder(r.Body).Decode(&cReq)
	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	if cReq.Name == "" {
		response.Status = "error"
		response.Message = "Company name must be supplied"
		SendJsonResponse(w, response)
		return
	}

	comp := models.Company{
		Name:        cReq.Name,
		Description: cReq.Description,
	}

	if err := b.CHandler.companyRepo.CreateCompany(comp); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
