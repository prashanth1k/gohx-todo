// handler.go

package main

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(ctx *Ctx, w http.ResponseWriter, r *http.Request) {
	log.Default().Println("IndexHandler")
	todoVal := map[string][]Todo{"Todos": ctx.Todos}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, todoVal)
}
