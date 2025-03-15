package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Connect to db
	// initDB()

	http.HandleFunc("/tasks", getTasksHandler)

	fmt.Println(" TODO App -- Using Golang ")
	port := ":8083"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
