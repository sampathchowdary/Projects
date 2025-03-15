package main

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode((response{
		Status:  statusCode,
		Message: message,
	}))
}

func sendSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode((response{
		Status: statusCode,
		Data:   data,
	}))
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasks()
	if err != nil {
		sendErrorResponse(w, 400, err.Error())
	}
	sendSuccessResponse(w, 200, tasks)
}
