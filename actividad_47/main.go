/*Cargar una matriz B de 6x6, ordenar cada columna en forma ascendente. Imprimir la matriz resultante.*/

package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

func ordenarMatriz(A *[filas][columnas]int) {
	for j := 0; j < columnas; j++ {
		for i := 0; i < filas-1; i++ {
			for k := 0; k < filas-i-1; k++ {
				A[k][j], A[k+1][j] = A[k+1][j], A[k][j]
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

/*




 */
