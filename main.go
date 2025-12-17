package main

import (
	"fmt"
	"net/http"

	"github.com/PittsGitHub/lenslocked/controllers"
	"github.com/PittsGitHub/lenslocked/models"
	"github.com/PittsGitHub/lenslocked/templates"
	"github.com/PittsGitHub/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}

	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)

	r.Post("/signup", usersC.Create)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
