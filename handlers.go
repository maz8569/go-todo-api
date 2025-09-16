package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func listHandler(w http.ResponseWriter, _ *http.Request) {
	tasks := loadTasksServer(w)
	if tasks == nil {
		return
	}
	//json.NewEncoder(w).Encode(tasks)

	for i, t := range tasks {
		status := 'X'
		if t.Done {
			status = 'V'
		}
		fmt.Fprintf(w, "%d) [%c] %s (created %s) [id: %d]\n", i+1, status, t.Title, t.Date.Format("Mon, 02.01.2006 (15-04)"), t.ID)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", 400)
	}

	tasks := loadTasksServer(w)
	if tasks == nil {
		return
	}

	newTask := Task{
		ID:    tasks[len(tasks)-1].ID + 1,
		Title: input.Title,
		Done:  false,
		Date:  time.Now(),
	}

	tasks = append(tasks, newTask)

	if !saveTasksServer(w, tasks) {
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func doneHandler(w http.ResponseWriter, r *http.Request) {

	tasks := loadTasksServer(w)
	if tasks == nil {
		return
	}

	id := findTaskByStringID(w, r, tasks)

	if id == -1 {
		return
	}

	tasks[id].Done = true

	if !saveTasksServer(w, tasks) {
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task \"%s\" marked as done.", tasks[id].Title)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	tasks := loadTasksServer(w)
	if tasks == nil {
		return
	}

	id := findTaskByStringID(w, r, tasks)

	if id == -1 {
		return
	}

	deletedTaskName := tasks[id].Title
	tasks = append(tasks[:id], tasks[id+1:]...)

	if !saveTasksServer(w, tasks) {
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task \"%s\" deleted.", deletedTaskName)

}

func loadTasksServer(w http.ResponseWriter) []Task {
	tasks, err := loadTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks.", 500)
		return nil
	}
	return tasks
}

func saveTasksServer(w http.ResponseWriter, tasks []Task) bool {
	if err := saveTasks(tasks); err != nil {
		http.Error(w, "failed to save task", http.StatusInternalServerError)
		return false
	}
	return true
}

func findTaskByStringID(w http.ResponseWriter, r *http.Request, tasks []Task) int {
	param := mux.Vars(r)
	idStr := param["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return -1
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			found = true
			id = i
			break
		}
	}

	if !found {
		http.Error(w, "task with this id "+idStr+" not found", http.StatusNotFound)
	}

	return id
}
