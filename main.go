package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
	Name  string
	Job   string
	Email string
	LinkedIn string
	Github string
	Body  string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error to load Template",http.StatusInternalServerError)
			return
		}

		data := PageData{

			Title: "Resume",
			Name: "Alexandre Julião",
			Job: "Software Developer",
			Email: "ajuliao-@",
			LinkedIn: "https://www.linkedin.com/in/alexandre-juliao/",
			Github: "",
			Body: "Welcome to my personal website",
		}
		tmpl.Execute(w, data)
	})
	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
