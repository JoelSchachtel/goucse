/*Cargar una matriz A de 5x5, ordenar cada fila en forma ascendente. Imprimir la matriz resultante.*/

package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

func ordenarMatriz(A *[filas][columnas]int) {

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas-1; j++ {
			for k := 0; k < columnas-j-1; k++ {
				if A[i][k] > A[i][k+1] {
					A[i][k], A[i][k+1] = A[i][k+1], A[i][k]
				}
			}
		}
	}

}

func main() {

	var (
		A [filas][columnas]int
	)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Println("Ingrese un N°: ")
			fmt.Scanln(&A[j][i])
		}
	}

	ordenarMatriz(&A)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Print(A[i][j], " ")
		}
		fmt.Println("")
	}

}
