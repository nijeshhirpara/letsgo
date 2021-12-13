package api

import (
	"encoding/json"
	"letsgo/models"
	"log"
	"net/http"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	CHandler *CompanyHandler
	THandler *TeamHandler
	UHandler *UserHandler
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(companyRepo models.CompanyRepository, teamRepo models.TeamRepository, userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		CHandler: &CompanyHandler{
			companyRepo: companyRepo,
		},
		THandler: &TeamHandler{
			teamRepo: teamRepo,
		},
		UHandler: &UserHandler{
			userRepo: userRepo,
		},
	}
}

type SuccessResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Status string   `json:"status"`
	Msg    []string `json:"msg"`
}

func SendJsonResponse(w http.ResponseWriter, res interface{}) {
	jsonContent, jsonError := json.MarshalIndent(res, "", "	")
	if jsonError != nil {
		log.Println("HandleSettings: Error in JSON marshal", jsonError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonContent)
}
