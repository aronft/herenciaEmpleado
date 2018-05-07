package main

import (
	"fmt"
	"log"
)

//Hijo es una estructura hijo de Persona
type Hijo struct {
	Persona
	Discapacidad bool
	Detalles     string
}

//CrearHijo Crea un hijo del empleado
func CrearHijo(h Hijo) int {

	hijoID, valid := h.CrearPersona()
	if valid == true {
		return hijoID
	}
	q := `INSERT INTO 
			hijos(hijo_id, discapacidad, detalles)
			VALUES ($1, $2, $3);`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(hijoID, h.Discapacidad, h.Detalles)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta: %s\n", err)
	}

	// i, _ := r.RowsAffected()
	// if i != 1 {
	// 	log.Fatal("Error: Se esperbaa una fila afectda")
	// }

	return hijoID
}

//ActualizarHijo actualiza los datos de los hijos
func ActualizarHijo(h Hijo) {

	h.ActualizarPersona()

	q := `UPDATE hijos SET discapacidad = $1, detalles $2 WHERE hijos_id = $3`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al prepapar la consulta: %s\n", err)
	}

	r, err := stmt.Exec(h.Discapacidad, h.Detalles, h.ID)
	if err != nil {
		fmt.Printf("Eror al ejecutar la consulta: %s\n", err)
	}

	i, err := r.RowsAffected()
	if i != 1 {
		fmt.Printf("Se esperaba almenos 1 fila afectada: %s", err)
	}
}
