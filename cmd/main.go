package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VillaJoseLuisJr/practica2/internal/config"
	"github.com/VillaJoseLuisJr/practica2/internal/controllers"
)

// var nombre = "Peronista"

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

	// Llamo a la función SetupRoutes del archivo rutas.go dentro de la carpeta config para que maneje las url
	controllers.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
