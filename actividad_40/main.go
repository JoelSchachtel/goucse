// Cargar una matriz A de 6x6, generar una nueva matriz de igual dimensión a la dada bajo las siguientes
// condiciones:
// a) si el elemento es negativo colocar un 0 en la nueva matriz.
// b) si el elemento es cero colocar un 1 en la nueva matriz.
// c) si el elemento es positivo colocar un 2 en la nueva matriz.
// d) Generar un procedimiento que imprima la matriz pasada por parámetro. Utilizar el procedimiento para
// imprimir las dos matrices

package main

import "fmt"

const (
	columna = 2
	fila    = 2
)

var (
	A [fila][columna]int
)

func matrizNueva(A [fila][columna]int) [fila][columna]int {
	var B [fila][columna]int

	for i := 0; i < fila; i++ {
		for j := 0; j < columna; j++ {
			if A[i][j] > 0 {
				B[i][j] = 2
			} else if A[i][j] == 0 {
				B[i][j] = 1
			} else {
				B[i][j] = 0
			}
		}
	}

	return B
}

func main() {

	var (
		A [fila][columna]int
		B [fila][columna]int
	)

	for i := 0; i < fila; i++ {
		for j := 0; j < columna; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&A[i][j])
		}
	}

	B = matrizNueva(A)

	for i := 0; i < fila; i++ {
		for j := 0; j < columna; j++ {
			fmt.Printf("%d", B[i][j])
		}
		fmt.Println("")
	}
}
