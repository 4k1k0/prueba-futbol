package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/4k1k0/prueba-futbol/data"
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

func nombreArchivo(archivo string) string {
	arreglo := strings.Split(archivo, "/")
	archivo = arreglo[len(arreglo)-1]
	arreglo = strings.Split(archivo, ".")
	archivo = arreglo[0]
	return archivo
}

func guardar(resultados []data.Resultado) (string, error) {
	data := "["
	for i := 0; i < len(resultados); i++ {
		out, err := json.MarshalIndent(resultados[i], "  ", "  ")
		if err != nil {
			log.Fatal(err)
		}
		cadena := string(out)
		if i == len(resultados) || i == 0 {
			data += "\n  " + cadena
		} else {
			data += ",\n  " + cadena
		}
	}
	data += "\n]\n"
	data2 := []byte(data)
	nombre := nombreArchivo(*equipo)
	archivo := "/tmp/" + nombre + ".log"
	err := ioutil.WriteFile(archivo, data2, 0644)
	msj := "Archivo generado en " + archivo
	return msj, err
}
