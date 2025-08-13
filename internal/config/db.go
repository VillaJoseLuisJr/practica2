package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	// Configuración de la conexión
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root123",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "practica2_go",
		AllowNativePasswords: true,
	}

	// Handle para la base de datos

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conexión con la base de datos exitosa!")
}
