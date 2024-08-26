/*
Se tiene las calificaciones de 6 asignaturas (enumeradas del 1 al 6) correspondientes a 20 alumnos
(enumerados del 1 al 20):
	a) Cargar una matriz AL con los datos de las calificaciones (cada fila representa un alumno y cada columna
	una materia).
	b) Crear una función que calcule y devuelva como resultado un vector con el promedio de cada asignatura.
	c) Crear una función que calcule y devuelva como resultado un vector con el promedio de cada estudiante.
	d) Generar una función que calcule y devuelva como resultado la cantidad de materias que aprobó y
	desaprobó un estudiante “x” pasado como parámetro.
*/

package main

import (
	"fmt"
)

const (
	filas    = 20
	columnas = 6
)

var (
	AL [filas][columnas]int
)

func promedioMaterias(AL [filas][columnas]int) [columnas]float64 {
	var promedios [columnas]float64

	for j := 0; j < columnas; j++ {
		suma := 0
		for i := 0; i < filas; i++ {
			suma += AL[i][j]
		}
		promedios[j] = float64(suma) / float64(filas)
	}

	return promedios
}

func promedioAlumnos(AL [filas][columnas]int) [filas]float64 {
	var promedios [filas]float64

	for i := 0; i < filas; i++ {
		suma := 0
		for j := 0; j < columnas; j++ {
			suma += AL[i][j]
		}
		promedios[i] = float64(suma) / float64(columnas)
	}

	return promedios
}

func aprobadasDesaprobadas(AL [filas][columnas]int, estudiante int) (int, int) {
	aprobadas := 0
	desaprobadas := 0

	for i := 0; i < columnas; i++ {
		if AL[estudiante][i] >= 6 {
			aprobadas++
		} else {
			desaprobadas++
		}
	}

	return aprobadas, desaprobadas
}

func main() {

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Printf("Ingrese la nota de la materia %d del alumno N° %d: ", j+1, i+1)
			fmt.Scanln(&AL[i][j])
		}
	}

	promedioMat := promedioMaterias(AL)
	fmt.Println("Promedio de cada materia:", promedioMat)

	promedioAl := promedioAlumnos(AL)
	fmt.Println("Promedio de cada alumno:", promedioAl)

	var estudiante int

	fmt.Print("Ingrese el N° de estudiante para calcular sus materias (1-20): ")
	fmt.Scanln(&estudiante)
	for estudiante != 0 {
		if estudiante > 20 {
			fmt.Println("Ingrese un estudiante válido: 1 - 20")
		} else if estudiante < 0 {
			fmt.Println("Ingrese un estudiante válido: 1 - 20")
		} else {
			aprobadas, desaprobadas := aprobadasDesaprobadas(AL, estudiante-1)
			fmt.Printf("El estudiante %d aprobó %d materias y desaprobó %d materias\n", estudiante, aprobadas, desaprobadas)
		}
		fmt.Print("Ingrese el N° de estudiante para calcular sus materias (1-20): ")
		fmt.Scanln(&estudiante)
	}
}
