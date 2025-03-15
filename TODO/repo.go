package main

func getTasks() ([]Task, error) {

	tasks := []Task{
		{
			ID:        1,
			Title:     "Learn Go Programming",
			Notes:     "Complete the Go to-do list project",
			Completed: "No",
			DueDate:   "2025-03-20",
		},
		{
			ID:        2,
			Title:     "Write a Blog Post",
			Notes:     "Write a blog post about Go language",
			Completed: "Yes",
			DueDate:   "2025-03-15",
		},
	}
	return tasks, nil
}
