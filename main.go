package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"

	"di-test/templates"
)

func main() {
	component := templates.Hello("world")
	http.Handle("/", templ.Handler(component))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf(err.Error())
	}
}
