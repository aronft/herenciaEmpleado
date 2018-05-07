package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "aster123"
	dbname   = "herencia"
)

//GetConection Realiza una conexion a la base de datos
func GetConection() *sql.DB {
	// c := GetConfiguration()
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s, sslmode=disable", host, port, user, password, dbname)
	dsn := "postgres://postgres:aster123@127.0.0.1:5432/herenciaempleado?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %s", err)
	}
	// defer db.Close()

	return db

}
