package main

import "fmt"

func main() {
	var (
		n1 int
		n2 int
		n3 int
	)
	fmt.Println("Ingrese un N°:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)

	if n1 > n2 && n1 > n3 {
		fmt.Println("El número", n1, "es el mayor de los 3.")
	} else if n2 > n3 {
		fmt.Println("El número", n2, "es el mayor de los 3.")
	} else {
		fmt.Println("El número", n3, "es el mayor de los 3.")
	}
}
