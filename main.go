package main

import (
	"fmt"
	"net/http"

	"github.com/PittsGitHub/lenslocked/controllers"
	"github.com/PittsGitHub/lenslocked/models"
	"github.com/PittsGitHub/lenslocked/templates"
	"github.com/PittsGitHub/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	// Setup a database connection
	pgCFG := models.DefaultPostgresConfig()
	db, err := models.Open(pgCFG)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userServices := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userServices,
	}

	// User Scoped Pages
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	r.Get("/users/me", usersC.CurrentUser)

	// User Agnostic Pages
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

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix this before deploying
		csrf.Secure(false),
		// Note: This is required if using v1.7.3+
		// due to a breaking change made to fix a
		// security issue.
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	fmt.Println("Starting the server on :3000...")

	handler := logRequests(r) // your logger first
	handler = csrfMw(handler) // then CSRF protection

	http.ListenAndServe(":3000", handler)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// RemoteAddr is usually "IP:port"
		fmt.Printf("REQ ip=%s method=%s path=%s ua=%q\n",
			r.RemoteAddr, r.Method, r.URL.Path, r.UserAgent(),
		)
		next.ServeHTTP(w, r)
	})
}
