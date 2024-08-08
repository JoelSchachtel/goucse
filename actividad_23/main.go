/*
Cargar un vector A de 10 elementos enteros. Ingresar un número entero N y contar la cantidad de veces
que aparece dicho número en el vector. Imprimir el resultado.
*/
package main

import "fmt"

var (
	numeros [10]int
	n       int
	nb      int
	contb   int
)

func main() {
	fmt.Println("Ingrese un N° entero:")
	fmt.Scanln(&n)

	for i := 0; i < len(numeros); i++ {
		numeros[i] = n

		if i < len(numeros)-1 { // Evitar pedir más números después del último
			fmt.Println("Ingrese un N° Entero: ")
			fmt.Scanln(&n)
		}
	}

	fmt.Println("Ingrese el N° a buscar")
	fmt.Scanln(&nb)
	for i := 0; i < len(numeros); i++ {
		if numeros[i] == nb {
			contb = contb + 1
		}
	}

	fmt.Printf("El numero %d fue contado: %d veces.\n", nb, contb)
	fmt.Println(numeros)
}
