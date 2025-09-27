package controllers

import (
	"net/http"
)

func SetupRoutes() {
	Tpl, _ = Tpl.ParseGlob("../web/templates/*.html")
	fs := http.FileServer(http.Dir("../web/static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", MostrarHTML)
	http.HandleFunc("/formulario", MostrarFormulario)
	http.HandleFunc("/guardar", GuardarProducto)
}
