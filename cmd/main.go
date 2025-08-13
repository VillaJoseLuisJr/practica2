package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/VillaJoseLuisJr/practica2/internal/config"
)

var tpl *template.Template

// var nombre = "Peronista"

func MostrarHTML(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "formulario.html", nil)
}

func main() {

	// Código de prueba
	rows, err := config.DB.Query("SELECT id, nombre, descripcion FROM producto")
	if err != nil {
		fmt.Println("Error en la consulta:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre, descripcion string
		if err := rows.Scan(&id, &nombre, &descripcion); err != nil {
			fmt.Println("Error leyendo fila:", err)
			return
		}
		fmt.Printf("%d: %s - %s\n", id, nombre, descripcion)
	}
	// Código de prueba

	tpl, _ = tpl.ParseGlob("../web/templates/*.html")
	fs := http.FileServer(http.Dir("../web/static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", MostrarHTML)
	http.HandleFunc("/formulario", MostrarFormulario)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
