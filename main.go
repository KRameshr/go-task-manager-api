package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Task represents a single to-do item
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// In-memory list of tasks
var tasks []Task

// GET /tasks - Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// POST /tasks - Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

// PUT /tasks/{id} - Update a task
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...) // remove old task
			var updated Task
			_ = json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = params["id"]
			tasks = append(tasks, updated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// DELETE /tasks/{id} - Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	fmt.Println("🚀 Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
