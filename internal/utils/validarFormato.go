package utils

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

func ValidarFormatoImagen(file multipart.File) (string, error) {
	//Leer los primeros 512 bytes del archivo
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "Error al leer la cabecera del archivo: ", err
	}

	//Detectar el tipo
	contentType := http.DetectContentType(buffer)

	//Resetear el puntero para poder volver a leer el archivo
	file.Seek(0, 0)

	switch contentType {
	case "image/jpeg", "image/png", "image/gif":
		log.Println("Tipo de formato detectado: ")
		return contentType, nil
	default:
		return "", fmt.Errorf("formato no permitido: %s", contentType)
	}

}
