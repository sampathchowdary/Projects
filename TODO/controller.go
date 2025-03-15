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

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := createTask(task.Title, task.Notes, task.DueDate); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(w, 200, "Task created successfully")
}

func updateTasksStatusHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID        int  `json:"id"`
		Completed bool `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := updateTaskStatus(request.ID, request.Completed); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(w, 200, "Task status updated successfully")
}
