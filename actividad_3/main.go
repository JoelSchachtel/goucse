package main

import "fmt"

var (
	n1 int
	n2 int
)

func main() {
	fmt.Println("Ingrese dos números:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	if n1 > n2 {
		fmt.Println("El primer número es mayor que el segundo, el número es:", n1)
	} else {
		fmt.Println("El segundo número es mayor que el primero, el número es:", n2)
	}
}
