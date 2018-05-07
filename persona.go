package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//Persona es el modelo padre de todos
type Persona struct {
	ID        int
	Dni       string
	Nombre    string
	Apellidos string
	FechaNac  time.Time
}

//CrearPersona crea personas para empelados e hijos
func (p *Persona) CrearPersona() (int, bool) {
	var id int
	var valid bool
	validateq := `SELECT persona_id, EXISTS(SELECT 1 FROM personas WHERE dni = $1) FROM personas WHERE dni = $1;

	`

	q := `INSERT INTO
	personas (dni, nombres, apellidos, fecha_nac)
	VALUES ($1, $2, $3, $4) RETURNING persona_id;`

	db := GetConection()
	defer db.Close()

	stmt2, err := db.Prepare(validateq)
	if err != nil {
		log.Fatalf("Erorr al validar: %s", err)
	}
	defer stmt2.Close()

	err = stmt2.QueryRow(p.Dni).Scan(&id, &valid)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No existe nadie con ese ID")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Persona ID es %d\n", id)
	}

	if id == 0 {
		stmt, err := db.Prepare(q)
		if err != nil {
			log.Fatalf("Erorr al preparar la consulta: %s", err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(p.Dni, p.Nombre, p.Apellidos, p.FechaNac).Scan(&id)
		if err != nil {
			log.Fatalf("erroa la ejecutar la consulta: %s", err)
		}
	}

	// i, _ := r.RowsAffected()
	// if i != 1 {
	// 	log.Fatalf("ninguna fila afectada: %s", err)
	// }
	return id, valid
}

//ActualizarPersona Actualiza los datos de las personas
func (p *Persona) ActualizarPersona() {
	q := `UPDATE personas SET dni = $1, nombres = $2, apellidos = $3, fecha_nac = $4 WHERE persona_id = $5`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al preparar la consulta: %s'\n", err)
	}

	r, err := stmt.Exec(p.Dni, p.Nombre, p.Apellidos, p.FechaNac, p.ID)
	if err != nil {
		fmt.Printf("Error al ejecutar la consulta: %s\n", err)
	}

	i, err := r.RowsAffected()
	if i != 1 {
		fmt.Printf("Error: Se esperaba una fila efctada: %", err)
	}
}
