package main

import "fmt"

func main() {
	var (
		numero  [10]int
		digito5 int
	)

	fmt.Println("Ingrese 10 números:")

	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numero[i])
		if numero[i]%5 == 0 {
			digito5 += 1
		}
	}

	fmt.Printf("La cantidad de números mútiplos de 5 es %d", digito5)
}
