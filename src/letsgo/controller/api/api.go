package api

import (
	"encoding/json"
	"log"
	"net/http"
)

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
