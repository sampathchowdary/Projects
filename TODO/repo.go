package main

func getTasks() ([]Task, error) {

	rows, err := db.Query(`
		SELECT id, title, notes, completed, due_date, priority
		FROM tasks 
		ORDER BY completed, priority, due_date, id
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // close connection at end of getTasks func

	tasks := []Task{}
	for rows.Next() {
		var task Task

		// Scan current row and insert data in each of these pointers
		if err := rows.Scan(&task.ID, &task.Title, &task.Notes, &task.Completed, &task.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func createTask(title, notes, dueDate string) error {
	_, err := db.Exec(`
		INSERT INTO tasks (title, notes, completed, due_date) VALUES (?, ?, ?, ?)`,
		title, notes, false, dueDate,
	)
	return err
}

func updateTaskStatus(id int, completed bool) error {
	_, err := db.Exec("UPDATE tasks SET completed = ? WHERE id = ?", completed, id)
	return err
}
