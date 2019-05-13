package main

import "flag"

var equipo = flag.String("equipo", "resuelve.json", "El equipo a analizar")
var niveles = flag.String("niveles", "niveles.json", "Los niveles con los cuales calcular")

func main() {
	flag.Parse()
	abrir(*equipo, *niveles)
}
