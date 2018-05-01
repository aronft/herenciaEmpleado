package models

//Persona es el modelo padre de todos
type Persona struct {
	id     int    `json:"id"`
	nombre string `json:"nombre"`
}
