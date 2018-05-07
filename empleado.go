package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

// Empleado es un hijo de la estructura Persona
type Empleado struct {
	Persona
	Sueldo float64
}

//CrearEmpleado Registra un empleado
func CrearEmpleado(e Empleado) int {
	empleadoID, valid := e.CrearPersona()
	if valid == true {
		return empleadoID
	}
	q := `INSERT INTO 
			empleados (empleado_id, sueldo)
			VALUES ($1, $2)	`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Erro al prepara la consulta empleado: %s\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(empleadoID, e.Sueldo)
	if err != nil {
		fmt.Printf("Erro al ejecutar la consulta empleado: %s\n", err)
		// return empleadoID
	}

	// i, _ := r.RowsAffected()
	// if i != 1 {
	// 	log.Fatal("Error:Se esperaba 1 fila afectada")
	// }

	return empleadoID
}

//ActualizarEmpleado actualiza y manda a actualizar los datos de un empleado
func ActualizarEmpleado(e Empleado) {
	e.ActualizarPersona()

	q := `UPDATE empleados SET sueldo = $1 WHERE empleado_id = $2`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Erorr al preparar la consulta: %s\n", err)
	}

	r, err := stmt.Exec(e.Sueldo, e.Persona.ID)
	if err != nil {
		fmt.Printf("Error al ejecutar la consulta: %s", err)
	}

	i, err := r.RowsAffected()
	if i != 1 {
		fmt.Printf("Se esperaba al menos una fila afectada: %s", err)
	}

}

//EliminarEmpleado elimina todos los empleados
func EliminarEmpleado(EmpleadoID []int) {
	q := `DELETE FROM personas WHERE persona_id = $1`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al preparala consulta %s\n", err)
	}

	for _, id := range EmpleadoID {
		_, err := stmt.Exec(id)
		if err != nil {
			fmt.Printf("Eror al ejecutar la consulta: %s", err)
		}
	}
}
