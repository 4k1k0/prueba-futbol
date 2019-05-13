package main

import (
	"fmt"

	data "./data"
)

func calculo(jugadores []data.Jugador, niveles data.Niveles) {
	totalGolesEquipo, metaGolesEquipo := calculaGoles(jugadores, niveles)
	porcentajeGolesEquipo := porcentaje(totalGolesEquipo, metaGolesEquipo)
	fmt.Println(porcentajeGolesEquipo)
	for _, j := range jugadores {
		nivel := asignaNivel(j, niveles)
		porcentajeGolesIndividual := porcentaje(j.Goles, nivel)
		fmt.Println(porcentajeGolesIndividual)
	}
}

func calculaGoles(jugadores []data.Jugador, n data.Niveles) (uint64, uint64) {
	metaGolesEquipo := n.NivelA + n.NivelB + n.NivelC + n.NivelCuauh
	var totalGolesEquipo uint64
	for _, j := range jugadores {
		totalGolesEquipo += j.Goles
	}
	return metaGolesEquipo, totalGolesEquipo
}

func porcentaje(x, total uint64) float64 {
	return float64((x * 100) / total)
}

func asignaNivel(jugador data.Jugador, niveles data.Niveles) uint64 {
	switch jugador.Nivel {
	case "A":
		return niveles.NivelA
	case "B":
		return niveles.NivelB
	case "C":
		return niveles.NivelC
	case "Cuauh":
		return niveles.NivelCuauh
	}
	return 0
}
