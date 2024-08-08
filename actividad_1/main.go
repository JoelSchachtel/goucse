package main

import "fmt"

var (
	n1 int
	n2 int
	n3 int
)

func main() {
	var (
		mult  int
		resta int
		divi  int
	)
	fmt.Println("Ingrese 3 números")
	fmt.Scanln(&n1) // El & es solo para cuando el usuario del programa ingresa una variable
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	resta = n1 - n2
	fmt.Println("La diferencia del primero con el segundo es:", resta)
	mult = n2 * n3
	fmt.Println("La multiplicación de los dos ultimos es:", mult)
	divi = n1 / n3
	fmt.Println("La división del primero por el tercero es:", divi)
}
