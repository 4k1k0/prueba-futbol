package main

import (
	"fmt"

	"github.com/4k1k0/prueba-futbol/data"
)

func calculo(jugadores []data.Jugador, niveles data.Niveles) {
	totalGolesEquipo, metaGolesEquipo := calculaGoles(jugadores, niveles)
	porcentajeGolesEquipo := porcentaje(totalGolesEquipo, metaGolesEquipo)
	var resultados []data.Resultado
	for _, j := range jugadores {
		nivel := asignaNivel(j, niveles)
		porcentajeGolesIndividual := porcentaje(j.Goles, nivel)
		bono := calculaBono(j.Bono, porcentajeGolesEquipo, porcentajeGolesIndividual)
		sueldoTotal := calculaSueldoTotal(j.Sueldo, bono)
		resultado := generarResultado(j, nivel, bono, sueldoTotal)
		resultados = append(resultados, resultado)
	}
	msj, err := guardar(resultados)
	checkError(err)
	fmt.Println(msj)
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
	return (float64(x) * 100.00) / float64(total)
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

func calculaBono(valor uint64, porcentajeGolesEquipo, porcentajeGolesIndividual float64) float64 {
	if porcentajeGolesEquipo > 100.00 {
		porcentajeGolesEquipo = 100.00
	}
	if porcentajeGolesIndividual > 100.00 {
		porcentajeGolesIndividual = 100.00
	}
	porcentajeGolesEquipo = porcentajeGolesEquipo / 2
	porcentajeGolesIndividual = porcentajeGolesIndividual / 2
	porcentajeTotal := porcentajeGolesEquipo + porcentajeGolesIndividual
	bono := (float64(valor) * porcentajeTotal) / 100.00
	return bono
}

func calculaSueldoTotal(sueldo uint64, bono float64) float64 {
	return float64(sueldo) + bono
}

func generarResultado(j data.Jugador, golesMinimos uint64, bono, sueldoTotal float64) data.Resultado {
	resultado := data.Resultado{
		Nombre:         j.Nombre,
		GolesMinimos:   golesMinimos,
		Goles:          j.Goles,
		Sueldo:         j.Sueldo,
		Bono:           bono,
		SueldoCompleto: sueldoTotal,
		Equipo:         j.Equipo,
	}
	return resultado
}
