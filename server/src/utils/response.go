package utils

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Success bool `json:"success,omitempty"`
	Error bool `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func JSONResponse(w http.ResponseWriter, data interface{}, httpStatus int){
	setHeaders(w, httpStatus)
	json, _ := json.Marshal(data)
	w.Write(json)
}

func setHeaders(w http.ResponseWriter, httpStatus int){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpStatus)
}