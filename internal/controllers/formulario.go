package controllers

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/VillaJoseLuisJr/practica2/internal/config"
	"github.com/VillaJoseLuisJr/practica2/internal/models"
	"github.com/VillaJoseLuisJr/practica2/internal/utils"
)

func GuardarProducto(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos del formulario
	var nombreImagen string

	file, fileHeader, err := r.FormFile("imagen_input")
	if err != nil {
		http.Error(w, "Error al leer la imagen", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//validar tipo de archivo
	format, err := utils.ValidarFormatoImagen(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Imagen valida: ", fileHeader.Filename, "Formato: ", format)

	//Guardar imagen en la carpeta ./web/static/img/
	nombreArchivo := filepath.Base(fileHeader.Filename)
	ruta := filepath.Join("img", nombreArchivo)
	ruta = strings.ReplaceAll(ruta, "\\", "/")
	log.Println("Guardando imagen en: ", ruta)

	destino, err := os.Create(ruta)
	if err != nil {
		log.Println("Error al crear archivo en el servidor: ", err)
		http.Error(w, "No se pudo guardar la imagen", http.StatusInternalServerError)
		return
	}
	defer destino.Close()

	_, err = io.Copy(destino, file)
	if err != nil {
		log.Println("Error al copiar la imagen: ", err)
		http.Error(w, "Error al guardar la imagen", http.StatusInternalServerError)
		return
	}

	nombreImagen = ruta

	producto := models.Producto{
		Nombre:      r.FormValue("nombre_input"),
		Descripcion: html.EscapeString(r.FormValue("descripcion_input")),
		Imagen:      nombreImagen,
	}

	err = models.InsertProducto(config.DB, producto)
	if err != nil {
		http.Error(w, "Error al guardar el producto", http.StatusInternalServerError)
		return
	}

}
