/* Cargar una matriz A de 4x4, intercambiar los elementos correspondientes a las filas pares con los de las filas
impares (la fila 1 con la 2; la 3 con la 4). */

package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

var ()

func intercambiar(A *[filas][columnas]int) { //Se pasa la matriz como valor de referencia para que afecte a la matriz original
	fmt.Println(A)
	for i := 0; i < filas; i = +2 {
		for j := 0; j < columnas; j++ {
			A[i][j], A[i+1][j] = A[i+1][j], A[i][j]

		}
	}

}

func main() {

	var (
		A [filas][columnas]int
	)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&A[j][i])
		}
	}

	intercambiar(&A)
	fmt.Println(A)

}
