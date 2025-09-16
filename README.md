# go-todo-api (Go + Gorilla Mux)
A simple CRUD REST API for managing tasks, built in Go.
Data is persisted in a local JSON file for simplicity.

# features
GET /tasks – list all tasks
POST /tasks – add a new task
PUT /tasks/{id}/done – mark a task as done
DELETE /tasks/{id} – delete a task
