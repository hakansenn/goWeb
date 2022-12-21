package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles()
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
