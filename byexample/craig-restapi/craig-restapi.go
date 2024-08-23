package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Todo) String() string {
	return fmt.Sprintf("{id: %v, title: %v, completed? %v}", t.ID, t.Title, t.Completed)
}

var todos []Todo

func main() {
	elem := Todo{"1", "Do A", false}
	todos = append(todos, elem)
	http.HandleFunc("/todos", getTodos)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r)
	log.Printf("%v", r.Method)
	switch r.Method {
	case "GET":
		{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos)

		}
	case "POST":
		{
			var item Todo
			var bytes []byte

			//Why does this read 0 bytes??
			size, readerr := r.Body.Read(bytes[:])
			log.Printf("Error: %v, Size: %v", readerr, size)

			err := json.NewDecoder(r.Body).Decode(&item)
			log.Printf("Error: %v, Body: %v", err, item)
			todos = append(todos, item)
			log.Printf("Current: %v", todos)
		}
	default:
		{
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
