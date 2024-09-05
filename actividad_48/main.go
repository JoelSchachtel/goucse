/* Cargar una matriz A de 200x7, donde se almacenan las notas de 6 materias correspondientes a 200 alumnos.
La primera columna indica el nro. de legajo del alumno y las restantes las notas de cada materia. Se pide:
a) Generar una matriz M de orden Nx7, con los datos de los alumnos que tengan 60 puntos o más en todas las
materias.
b) Generar una matriz C de orden 200x8, similar a la matriz original colocando en la columna 8, el promedio
obtenido por el alumno.
c) Generar una función que ordene en forma ascendentemente por promedio la matriz C. La función deberá
recibir como parámetro por referencia la matriz a ordenar.
d) Imprimir la matriz C. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filas    = 200
	columnas = 7
)

func promedio(alumnos [filas][columnas]int) [filas][columnas + 1]int {

	var C [filas][columnas + 1]int

	for i := 0; i < filas; i++ {
		suma := 0
		for j := 1; j < columnas; j++ { // Arranca en 1 porque la primer columna son los legajos
			suma += alumnos[i][j]
			C[i][j] = alumnos[i][j]
		}
		promedio := suma / (columnas - 1) //Menos 1 porque una fila corresponde al legajo
		C[i][columnas] = promedio
		C[i][0] = alumnos[i][0] // Mantiene el legajo en la matriz C en la posicion 0
	}

	return C
}

func todasAprobadas(alumnos [filas][columnas]int) [][columnas]int {

	var M [][columnas]int

	for i := 0; i < filas; i++ {
		incluir := true
		for j := 0; j < columnas; j++ {
			if alumnos[i][j] < 60 {
				incluir = false
				break
			}
		}
		if incluir {
			M = append(M, alumnos[i])
		}

	}

	return M
}

func ordenarMatrizC(C *[filas][columnas + 1]int) {

	for i := 0; i < filas-1; i++ {
		for j := 0; j < filas-i-1; j++ {
			if C[j][columnas] > C[j+1][columnas] {
				for k := 0; k < columnas+1; k++ {
					C[j][k], C[j+1][k] = C[j+1][k], C[j][k]
				}
			}
		}
	}
}

func main() {
	var (
		nroRandom, legajo int
		alumnos           [filas][columnas]int
	)

	rand.Seed(time.Now().Unix())

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			if j == 0 {
				legajo = rand.Intn(9000) + 1000
				alumnos[i][j] = legajo
			} else {
				nroRandom = rand.Intn(101)
				alumnos[i][j] = nroRandom
			}

		}
	}

	M := todasAprobadas(alumnos)
	C := promedio(alumnos)
	ordenarMatrizC(&C)

	fmt.Print("Alumnos que aprobaron todas las materias: \n")
	for i := 0; i < len(M); i++ {
		for j := 0; j < columnas; j++ {
			fmt.Print(M[i][j], " ")
		}
		fmt.Println("")
	}

	fmt.Print("Lista de promedios ordenados: \n")
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas+1; j++ {
			fmt.Print(C[i][j], " ")
		}
		fmt.Println("")
	}

}
