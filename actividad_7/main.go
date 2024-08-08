package main

import "fmt"

func main() {
	var (
		n1 float64
		n2 float64
		n3 float64
	)

	fmt.Println("Ingrese 3 números enteros:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)

	if n1 <= n2 && n2 <= n3 {
		fmt.Println("Los números están ingresados están en orden creciente")
	} else {
		fmt.Println("Los números no están ingresados en orden creciente.")
	}

}
