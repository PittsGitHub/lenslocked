package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch email me at
	<a href="mailto:frontrowpittard@gmail.com">frontrowpittard@gmail.com</a>.`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<p>Open the door get on the floor, baby it's a </p>`)
	fmt.Fprint(w, `<h1>404 <span style="font-size: 0.5em;">not found</span></h1>`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	case "/":
		homeHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", router)
}
