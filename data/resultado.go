package data

// Resultado es la estructura que deberá regresar el programa
type Resultado struct {
	Nombre         string  `json:"nombre"`
	GolesMinimos   uint64  `json:"goles_minimos"`
	Goles          uint64  `json:"goles"`
	Sueldo         uint64  `json:"sueldo"`
	Bono           float64 `json:"bono"`
	SueldoCompleto float64 `json:"sueldo_completo"`
	Equipo         string  `json:"equipo"`
}
