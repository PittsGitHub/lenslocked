package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Name string
	Age  int
	Dog  Dog
}

type Dog struct {
	Name  string
	Breed string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml", nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml", nil)
}

func ownerHandler(w http.ResponseWriter, r *http.Request) {
	Dan := User{
		Name: "Dan",
		Age:  30,
		Dog: Dog{
			Name:  "Bolbus",
			Breed: "Hairy Borker 9000",
		},
	}
	executeTemplate(w, "templates/owner.gohtml", Dan)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	executeTemplate(w, "templates/notfound.gohtml", nil)
}

func main() {

	r := chi.NewRouter()
	r.With(middleware.Logger).Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/owner", ownerHandler)
	r.NotFound(notFoundHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

func executeTemplate(w http.ResponseWriter, fp string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}
