package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	host := chi.URLParam(r, "host")
	println(host)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	host := chi.URLParam(r, "host")
	println(host)
	fmt.Fprint(w, `
		<h1>FAQ</h1>
		<p>Frequently asked questions about Bilbo.</p>
		<p/>
		<ul>
			<li>
				<strong>Who is Bilbo?</strong><br>
				Bilbo is a small scruffy black haired jack-a-poo with a knack for getting into (and out of) trouble.
			</li>
			<p/>
			<li>
				<strong>What does Bilbo do all day?</strong><br>
				Mostly minds his own business, snacks frequently, and occasionally goes on unexpected walks.
			</li>
			<p/>
			<li>
				<strong>Does Bilbo like visitors?</strong><br>
				Yes. Other dogs. People. Even cats.
			</li>
			<p/>
			<li>
				<strong>How can I learn more about Bilbo?</strong><br>
				You can email Bilbo at 
				<a href="mailto:bilbo@borks.alot">bilbo@borks.alot</a>.
			</li>
		</ul>
	`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Contact page for id: %s", id)
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch email me at
	<a href="mailto:frontrowpittard@gmail.com">frontrowpittard@gmail.com</a>.`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	host := chi.URLParam(r, "host")
	println(host)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<p>Open the door get on the floor, baby it's a </p>`)
	fmt.Fprint(w, `<h1>404 <span style="font-size: 0.5em;">not found</span></h1>`)
}

func main() {
	r := chi.NewRouter()
	r.With(middleware.Logger).Get("/", homeHandler)
	r.With(contactLogMiddleware).Get("/contact/{id}", contactHandler)
	r.Get("/faq{host}", faqHandler)
	r.NotFound(notFoundHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}

func contactLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		log.Printf("contact hit with id=%s", id)
		next.ServeHTTP(w, r)
	})
}
