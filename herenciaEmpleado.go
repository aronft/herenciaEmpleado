package main

func main() {
	/*
		var eh EmpleadoHijo
		eh.empleado = Empleado{
			Persona: Persona{
				Dni:       "5656",
				Nombre:    "Jorge",
				Apellidos: "Flores",
			},
			Sueldo: 890,
		}
		eh.hijosList = []Hijo{
			{
				Persona: Persona{
					Dni:       "123",
					Nombre:    "Jorge",
					Apellidos: "Tapia",
				},
			},
			{
				Persona: Persona{
					Dni:       "222",
					Nombre:    "Maria",
					Apellidos: "Gutierrez",
				},
			},
			{
				Persona: Persona{
					Dni:       "123",
					Nombre:    "Pedro",
					Apellidos: "Cruz",
				},
			},
		}
	*/
	/* var HijosID []int
	var nuevoID, empleadoID int

	nuevoID = 33
	empleadoID = 40
	HijosID = append(HijosID, 35)

	ActualizarEmpleadoHijo(nuevoID, empleadoID, HijosID)
	*/
	//Actualizar empleado
	// var e Empleado
	// e = Empleado{
	// 	Persona: Persona{
	// 		ID:        40,
	// 		Dni:       "565656",
	// 		Nombre:    "Jorgito",
	// 		Apellidos: "Flores",
	// 	},
	// 	Sueldo: 1600,
	// }

	// ActualizarEmpleado(e)

	//Eliminar empleado

	var empleadosID []int

	empleadosID = append(empleadosID, 40)

	EliminarEmpleado(empleadosID)
	// err := CrearEmpleadoHijo(eh)

	// fmt.Println("Creado exitosamente")

}
