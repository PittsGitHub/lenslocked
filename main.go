package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my big air site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/big", handlerFunc2)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}
