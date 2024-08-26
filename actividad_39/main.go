// Cargar una matriz A de 5x3 e imprimir el menor elemento y la posición donde se encuentra (fila y columna).

package main

import "fmt"

var (
	matriz [5][3]int
)

func menorElemento(matriz [5][3]int) (int, int, int) {
	var (
		menorElemento = matriz[0][0]
		fila, columna int
	)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if matriz[i][j] < menorElemento {
				menorElemento = matriz[i][j]
				fila = i
				columna = j
			}
		}
	}

	return menorElemento, fila, columna
}

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&matriz[i][j])
		}
	}

	menor, fila, columna := menorElemento(matriz)

	fmt.Printf("El menor elemento es %d, ubicado en la matriz en la posición %d, %d\n", menor, fila, columna)
}
