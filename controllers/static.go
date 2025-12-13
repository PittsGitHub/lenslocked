package controllers

import (
	"net/http"

	"github.com/PittsGitHub/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Who is Bilbo?",
			Answer:   "Bilbo is a small scruffy black haired jack-a-poo with a knack for getting into (and out of) trouble.",
		},
		{
			Question: "What does Bilbo do all day?",
			Answer:   "Mostly minds his own business, snacks frequently, and occasionally goes on unexpected walks.",
		},
		{
			Question: "Does Bilbo like visitors?",
			Answer:   "Yes. Other dogs. People. Even cats.",
		},
		{
			Question: "How can I learn more about Bilbo?",
			Answer:   `You can email Bilbo at <a href="mailto:bilbo@borks.alot">bilbo@borks.alot</a>.`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}

}
