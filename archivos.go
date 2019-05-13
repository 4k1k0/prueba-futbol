package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	data "./data"
)

func abrir(equipo, niveles string) {
	err := abrirArchivo(equipo)
	checkError(err)
	err = abrirArchivo(niveles)
	checkError(err)
	equipoCadena := generaCadena(equipo)
	nivelesCadena := generaCadena(niveles)
	generaData(equipoCadena, nivelesCadena)
}

func abrirArchivo(archivo string) error {
	f, err := os.Open(archivo)
	defer f.Close()
	return err
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generaCadena(archivo string) string {
	b, err := ioutil.ReadFile(archivo)
	checkError(err)
	s := string(b)
	return s
}

func generaData(equipo, nivel string) {
	jugadores := generaJugadores(equipo)
	niveles := generaNiveles(nivel)
	calculo(jugadores, niveles)
}

func generaJugadores(archivo string) []data.Jugador {
	var jugadores []data.Jugador
	err := json.Unmarshal([]byte(archivo), &jugadores)
	checkError(err)
	return jugadores
}

func generaNiveles(archivo string) data.Niveles {
	var nivel data.Niveles
	err := json.Unmarshal([]byte(archivo), &nivel)
	checkError(err)
	return nivel
}
