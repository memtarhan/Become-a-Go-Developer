package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)
func loadTemplates() {
	templateNames := [5]string {"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles(name + ".html")
		if (err == nil) {
			templates[name] = t 
			fmt.Println("Loaded template", index, name)
		
		} else {
			panic(err)
		}
	}
}

func welComeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}
func main() {
	loadTemplates()

	http.HandleFunc("/", welComeHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":3000", nil)
	if (err != nil) {
		fmt.Println(err)
	}
}