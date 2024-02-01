/*
So far, we have learned principles of GoLang including method, struct, instantiation, pointers. (Recall
than Golang does not have class).

Now we are ready to step into one of the super important applications of this language, which is API
development. See the below example.

Step 1: installing a web platform for Golang: go get github.com/gorilla/mux

Step 2: write the below code for API development.
*/

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Here, we define a struct that represents a task
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"` // The format of JSON is used mainly for avoiding compatibility issues
	// in the web. All the web plugins correspond to this format, and it is safe.
}

// In-memory store for tasks (we might want to use a database in a real-world scenario)
// Often, databases compatible with Golang are non-relational, such as MongoDB.
// We will cover it in the upcoming weeks.
// For now, a simple array is enough
var tasks []Task

// Now we define a GetTaskHandler function to return the list of tasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Here, we define another function as CreateTaskHandler to add a new task to the list
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func example1ApiDev() {

	tasks = append(tasks, Task{ID: "1", Title: "Task 1"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2"})
	// First, we create a new router
	router := mux.NewRouter()
	// Defining routes
	router.HandleFunc("/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")

	// Now we start the server
	http.ListenAndServe(":8080", router)
}

/*
In the above code we have defined tha task but have not assigned any value to it. SO when we run in
on the localhost, it shows null. The fact is that the "tasks" slice is initially empty and when we access
http://localhost:8080/tasks, it returns an empty JSON array ([]). TO verify that our API is working correctly,
we can use a tool like "curl" or postman to make a POST request to add a task and then check for tasks
again.

Here's how we use curl in golang (the same thing in all other web dev tools)
*/
