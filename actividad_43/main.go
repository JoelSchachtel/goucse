/*Cargar una matriz A de 5x3 y un vector B de 10 elementos. Sumar los elementos de A, que no encuentren su
igual en B. Imprimir la suma obtenida.*/

package main

import "fmt"

const (
	filas    = 5
	columnas = 3
	indice   = 10
)

func comparador(A [filas][columnas]int, B [indice]int) int {

	suma := 0

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			for k := 0; k < indice; k++ {
				if A[i][j] != B[k] {
					suma += A[i][j]
				}
			}
		}
	}

	return suma
}

func main() {

	var (
		A [filas][columnas]int
		B [indice]int
	)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Printf("Ingrese un N° para cargar en la matríz en la posición [%d][%d]", i, j)
			fmt.Scanln(&A[i][j])
		}
	}

	for k := 0; k < indice; k++ {
		fmt.Printf("Ingrese un N° para cargar en el vector en la posición [%d]", k)
		fmt.Scanln(&B[k])
	}

	resultado := comparador(A, B)

	fmt.Println(resultado)

}
