// Cargar una matriz de 6x6 elementos y poner un 0 en donde encuentre un valor par.

package main

import "fmt"

var (
	matriz [2][2]int
	n      int
)

func main() {

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&n)
			if n%2 == 0 {
				matriz[i][j] = 0
			} else {
				matriz[i][j] = n
			}
		}
	}

	fmt.Println(matriz)
}
