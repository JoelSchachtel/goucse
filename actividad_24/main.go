/*Cargar un vector de 7 números enteros y sumar aquellos elementos que sean distintos de su índice.
Imprimir la suma.*/

package main

import "fmt"

var (
	vector   [7]int
	n        int
	contador int
)

func main() {
	fmt.Println("Ingrese un N° entero:")
	fmt.Scanln(&n)

	for i := 0; i < len(vector); i++ {
		vector[i] = n

		if n != i {
			contador = contador + n
		}

		if i < len(vector)-1 {
			fmt.Println("Ingrese un N° entero:")
			fmt.Scanln(&n)
		}
	}

	fmt.Printf("La suma de los números que no coincide con su índice es: %d", contador)
}
