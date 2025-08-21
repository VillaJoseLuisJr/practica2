package config

import (
	"net/http"

	"github.com/VillaJoseLuisJr/practica2/internal/controllers"
)

func SetupRoutes() {
	controllers.Tpl, _ = controllers.Tpl.ParseGlob("../web/templates/*.html")
	fs := http.FileServer(http.Dir("../web/static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", controllers.MostrarHTML)
	http.HandleFunc("/formulario", controllers.MostrarFormulario)
}
