package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

// TaskServer holds the server's state.
type TaskServer struct {
	// store or any fields related to your server's state
}

// NewTaskServer creates a new TaskServer.
func NewTaskServer() *TaskServer {
	return &TaskServer{}
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	// Define routes
	mux.HandleFunc("/task/", server.taskHandler)
	mux.HandleFunc("/tag/", server.tagHandler)
	mux.HandleFunc("/due/", server.dueHandler)

	// Start server
	log.Println("Server is running at localhost:8090")
	http.ListenAndServe("localhost:8090", mux)
}

func (ts *TaskServer) taskHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		ts.createTaskHandler(w, req)
	case "GET":
		idStr := path.Base(req.URL.Path)
		if idStr == "task" {
			ts.getAllTasksHandler(w, req)
		} else {
			ts.getTaskHandler(w, req, idStr)
		}
	case "DELETE":
		idStr := path.Base(req.URL.Path)
		if idStr == "task" {
			ts.deleteAllTasksHandler(w, req)
		} else {
			ts.deleteTaskHandler(w, req, idStr)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ts *TaskServer) getTaskHandler(w http.ResponseWriter, req *http.Request, idStr string) {
	log.Printf("Handling get task at %s\n", req.URL.Path)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Replace the following with actual logic to retrieve a task
	// task, err := ts.store.GetTask(id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// Mock task
	task := struct {
		ID int `json:"id"`
	}{ID: id}

	renderJSON(w, task)
}

// Assuming Task is a struct representing a task
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// ... other fields ...
}

// createTaskHandler handles the creation of a new task
func (ts *TaskServer) createTaskHandler(w http.ResponseWriter, req *http.Request) {
	// Read request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Unmarshal the task from request body
	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	// TODO: Add logic to store the task
	// e.g., ts.store.CreateTask(task)

	w.WriteHeader(http.StatusCreated)
	renderJSON(w, task)
}

// getAllTasksHandler handles the retrieval of all tasks
func (ts *TaskServer) getAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: Add logic to retrieve all tasks
	// e.g., tasks, err := ts.store.GetAllTasks()

	// Mock response
	tasks := []Task{
		{ID: 1, Name: "Task One"},
		{ID: 2, Name: "Task Two"},
	}

	renderJSON(w, tasks)
}

// deleteAllTasksHandler handles the deletion of all tasks
func (ts *TaskServer) deleteAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: Add logic to delete all tasks
	// e.g., err := ts.store.DeleteAllTasks()

	w.WriteHeader(http.StatusOK)
}

// deleteTaskHandler handles the deletion of a single task by ID
func (ts *TaskServer) deleteTaskHandler(w http.ResponseWriter, req *http.Request, idStr string) {
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }

	// TODO: Add logic to delete the task by ID
	// e.g., err := ts.store.DeleteTask(id)

	w.WriteHeader(http.StatusOK)
}

// tagHandler handles the retrieval of tasks by tag
func (ts *TaskServer) tagHandler(w http.ResponseWriter, req *http.Request) {
	// Extract tag from URL
	// tag := path.Base(req.URL.Path)

	// TODO: Add logic to retrieve tasks by tag
	// e.g., tasks, err := ts.store.GetTasksByTag(tag)

	// Mock response
	tasks := []Task{
		{ID: 1, Name: "Task One"},
		// Tasks with the specified tag
	}

	renderJSON(w, tasks)
}

// dueHandler handles the retrieval of tasks by due date
func (ts *TaskServer) dueHandler(w http.ResponseWriter, req *http.Request) {
	// Extract due date components from URL
	// Assuming URL like /due/2023/04/05
	// parts := strings.Split(req.URL.Path, "/")
	// if len(parts) < 5 {
	// 	http.Error(w, "Invalid date format", http.StatusBadRequest)
	// 	return
	// }
	// year, month, day := parts[2], parts[3], parts[4]

	// TODO: Add logic to retrieve tasks by due date
	// e.g., tasks, err := ts.store.GetTasksByDueDate(year, month, day)

	// Mock response
	tasks := []Task{
		{ID: 1, Name: "Task One"},
		// Tasks due on the specified date
	}

	renderJSON(w, tasks)
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
