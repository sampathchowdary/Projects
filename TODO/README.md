# Golang To-Do App

This is a simple To-Do application built using **Go (Golang)**. The app exposes a RESTful API that allows you to retrieve a list of tasks.

## Features

- **GET /tasks**: Retrieve a list of to-do tasks in JSON format.

### Example Response:

```json
{
  "status": 200,
  "data": [
    {
      "id": 1,
      "title": "Learn Go Programming",
      "notes": "Complete the Go to-do list project",
      "completed": "No",
      "due_date": "2025-03-20"
    },
    {
      "id": 2,
      "title": "Write a Blog Post",
      "notes": "Write a blog post about Go language",
      "completed": "Yes",
      "due_date": "2025-03-15"
    },
    {
      "id": 3,
      "title": "Read Go Documentation",
      "notes": "Go through official Go documentation to improve understanding",
      "completed": "No",
      "due_date": "2025-03-25"
    }
  ]
}
