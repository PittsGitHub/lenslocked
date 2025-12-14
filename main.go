package main

import (
	"fmt"
	"net/http"

	"github.com/PittsGitHub/lenslocked/controllers"
	"github.com/PittsGitHub/lenslocked/templates"
	"github.com/PittsGitHub/lenslocked/views"
	"github.com/go-chi/chi/v5"
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

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(
		views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Get("/owner", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "owner.gohtml", "tailwind.gohtml"))))

	r.NotFound(controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "notfound.gohtml", "tailwind.gohtml"))))

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
