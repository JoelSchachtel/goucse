/*Cargar un vector de 5 números enteros e imprimir la suma y el promedio de dichos elementos.*/

package main

import "fmt"

var (
	vector   [5]int
	n        int
	suma     int
	promedio int
)

func main() {
	fmt.Println("Ingrese un N° Entero: ")
	fmt.Scanln(&n)

	for i := 0; i < len(vector); i++ {
		vector[i] = n

		suma = suma + n

		if i < len(vector)-1 { // Evitar pedir más números después del último
			fmt.Println("Ingrese un N° Entero: ")
			fmt.Scanln(&n)
		}
	}

	promedio = suma / len(vector)
	fmt.Printf("El vector cargado es: %d\n", vector)
	fmt.Printf("La suma de los números ingresados es: %d\n", suma)
	fmt.Printf("El promedio de los números ingresados es: %d\n", promedio)
}
