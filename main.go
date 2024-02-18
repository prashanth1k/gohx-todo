package main

import (
	"log"
	"net/http"
)

type Todo struct {
	Id    string
	Title string
	Done  bool
}

type Ctx struct {
	Todos []Todo
}

func main() {
	ctx := Ctx{Todos: []Todo{{Id: "1", Title: "Task 11", Done: false}, {Id: "2", Title: "Task 21", Done: true}}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		IndexHandler(&ctx, w, r)
	})

	http.HandleFunc("/api/todo", func(w http.ResponseWriter, r *http.Request) {
		TodoHandler(&ctx, w, r)
	})

	log.Default().Println("Server started. Your app is on http://localhost:8080")
	log.Default().Fatal(http.ListenAndServe(":8080", nil))

}
