package api

import (
	"encoding/json"
	"letsgo/model"
	"net/http"
)

func ListCompanies(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data   []model.Company `json:"data"`
		Status string          `json:"status"`
	}

	var response Response
	response.Status = "success"
	response.Data = model.ListCompanies()

	SendJsonResponse(w, response)
	return
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message`
	}

	var response Response

	type CompanyReq struct {
		Name        string `json:"name"`
		Description string `json:description`
	}

	var cReq CompanyReq
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

	c := model.Company{
		Name:        cReq.Name,
		Description: cReq.Description,
	}

	if err := model.CreateCompany(c); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		SendJsonResponse(w, response)
		return
	}

	response.Status = "success"
	SendJsonResponse(w, response)
	return
}
