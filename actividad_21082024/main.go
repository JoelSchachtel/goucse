package main

import "fmt"

var (
	matriz [2][3]int
	a      int
)

func main() {

	for i := 0; i < 2; i++ { //Recorre las filas
		for j := 0; j < 3; j++ { //Recorre las columnas
			fmt.Println("Ingrese un dato")
			fmt.Scanln(&a)
			matriz[i][j] = a
		}
	}

	fmt.Println(matriz)
}
