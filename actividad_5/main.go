package main

import "fmt"

func main() {
	var n int
	var ndiv int

	fmt.Println("Ingrese un N° para determinar si es PAR o IMPAR:")
	fmt.Scanln(&n)

	ndiv = n % 2

	if ndiv == 0 {
		fmt.Println("El N° ingresado", n, ", es un número PAR")
	} else {
		fmt.Println("El N° ingresado", n, ", es un número IMPAR")
	}

}
