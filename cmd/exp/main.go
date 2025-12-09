package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {

	fmt.Println("Experimental Main")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Bobby Joe",
		Bio:  `<script>alert("kek get hacked");</script>`,
		Age:  33,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
