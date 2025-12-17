package controllers

import "net/http"

type TemplateRenderer interface {
	Execute(w http.ResponseWriter, data any)
}
