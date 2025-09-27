package models

import (
	"database/sql"
	"fmt"
)

type Producto struct {
	ID          int64
	Nombre      string
	Descripcion string
	Imagen      string
}

func InsertProducto(db *sql.DB, p Producto) error {
	query := "INSERT INTO producto (nombre, descripcion, imagen) VALUES (?, ?, ?)"
	_, err := db.Exec(query, p.Nombre, p.Descripcion, p.Imagen)
	if err != nil {
		return fmt.Errorf("error al insertar producto %w", err)
	}
	return nil
}
