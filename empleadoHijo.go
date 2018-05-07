package main

import (
	"fmt"
)

//EmpleadoHijo es los hijos  y los padres
type EmpleadoHijo struct {
	empleado  Empleado
	hijosList []Hijo
}

//CrearEmpleadoHijo crea los empleados e hijos y gurda los detalles
func CrearEmpleadoHijo(eh EmpleadoHijo) error {
	var hijoID int

	empleadoID := CrearEmpleado(eh.empleado)

	q := `INSERT INTO 
			empleados_hijos(hijo_id, empleado_id)
			VALUES ($1, $2);`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	for _, hijo := range eh.hijosList {
		hijoID = CrearHijo(hijo)
		_, err = stmt.Exec(hijoID, empleadoID)
		if err != nil {
			fmt.Printf("Error:%s\n", err)
		}
	}

	return nil
}

//ActualizarEmpleadoHijo Actualiza todos los datos
func ActualizarEmpleadoHijo(nuevoID int, EmpleadoID int, HijosID []int) error {

	q := `UPDATE empleados_hijos SET empleado_id = $1 WHERE empleado_id = $2 AND hijo_id = $3;`

	db := GetConection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Erro al preparar la consulta: %s\n", err)
	}

	for _, Hijo := range HijosID {
		_, err := stmt.Exec(nuevoID, EmpleadoID, Hijo)
		if err != nil {
			fmt.Printf("Erorr al ejecutar la consulta: %s\n", err)
		}
	}
	return nil
}
