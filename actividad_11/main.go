package main

import "fmt"

func main() {

	var (
		numero  [10]int
		sumapar int
		sumaimp int
	)

	fmt.Println("Ingrese 10 números:")

	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numero[i])
		if numero[i]%2 == 0 {
			sumapar += numero[i]
		} else {
			sumaimp += numero[i]
		}
	}

	fmt.Printf("La suma de los pares es: %d, y la suma de los inpares es: %d", sumapar, sumaimp)
}
