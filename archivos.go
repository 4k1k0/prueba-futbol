package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	fmt.Printf("%s...%s", equipo, nivel)
}
