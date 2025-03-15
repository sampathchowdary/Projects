package main

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Notes     string `json:"notes"`
	Completed string `json:"completed"`
	DueDate   string `json:"due_date"`
}
type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
