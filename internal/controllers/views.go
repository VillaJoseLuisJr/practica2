package controllers

import (
	"html/template"
	"net/http"
)

var Tpl *template.Template

func MostrarHTML(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "index.html", nil)
}

func MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "formulario.html", nil)
}
