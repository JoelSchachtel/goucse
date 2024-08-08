package main

import (
	"fmt"
)

func main() {
	var (
		nro1, nro2, sumita int
	)
	fmt.Println("Ingrese el primer N°")
	fmt.Scanln(&nro1)
	fmt.Println("Ingrese el segundo N°")
	fmt.Scanln(&nro2)

	sumita = suma(nro1, nro2)

	fmt.Println(sumita)
}

func suma(nro1, nro2 int) int {
	return nro1 + nro2
}
