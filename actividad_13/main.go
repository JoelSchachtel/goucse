package main

import "fmt"

func main() {
	var numero int
	var numay int

	fmt.Println("Ingrese 10 números distintos:")

	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numero)
		if i == 0 {
			numay = numero
			fmt.Printf("El número ingresado: %d, es el primero y el más grande ingresado.\n", numero)
		} else {
			if numero > numay {
				numay = numero
				fmt.Printf("El número ingresado: %d, es el mayor de los ingresados.\n", numero)
			} else {
				fmt.Printf("El número ingresado: %d, no es mayor que %d\n", numero, numay)
			}
		}
	}
}
