package data

// Jugador es la estructura para almacenar la información de un jugador individual
type Jugador struct {
	Nombre         string `json:"nombre"`
	Nivel          string `json:"nivel"`
	Goles          uint64 `json:"goles"`
	Sueldo         uint64 `json:"sueldo"`
	Bono           uint64 `json:"bono"`
	SueldoCompleto uint64 `json:"sueldo_completo"`
	Equipo         string `json:"equipo"`
}
