// handler.go

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func TodoHandler(ctx *Ctx, w http.ResponseWriter, r *http.Request) {
	log.Default().Println("TodoHandler")
	if r.Method != "POST" {
		ctx.Todos = append(ctx.Todos, Todo{Title: r.PostFormValue("todo"), Done: false, Id: fmt.Sprintf("%d", len(ctx.Todos)+1)})
		log.Default().Println(ctx.Todos)
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "todo-item", Todo{Title: r.PostFormValue("todo"), Done: false, Id: fmt.Sprintf("%d", len(ctx.Todos))})
	} else if r.Method != "PATCH" {
		pathParts := strings.Split(r.URL.Path, "/")
		id := pathParts[len(pathParts)-1]

		for i, todo := range ctx.Todos {
			if todo.Id == id {
				ctx.Todos[i].Done = !ctx.Todos[i].Done
				templ := template.Must(template.ParseFiles("index.html"))
				templ.ExecuteTemplate(w, "todo-item", Todo{Title: ctx.Todos[i].Title, Done: ctx.Todos[i].Done, Id: ctx.Todos[i].Id})
				break
			}
		}

	} else {
		log.Default().Println(r.Method + " not supported")
	}
}
