package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var nombre = "Peron"

func MostrarHTML(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nombre)
}

func MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "formulario.html", nil)
}

func main() {
	tpl, _ = tpl.ParseGlob("../web/templates/*.html")
	fs := http.FileServer(http.Dir("../web/static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", MostrarHTML)
	http.HandleFunc("/formulario", MostrarFormulario)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
