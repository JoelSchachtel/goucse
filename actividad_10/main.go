package main

import "fmt"

var (
	numbers [8]float64
	sum     float64
)

func main() {

	fmt.Println("Ingrese 8 números: ")

	for i := 0; i < 8; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numbers[i])
		sum += numbers[i]
	}

	promedio := sum / 8
	fmt.Printf("El proomedio es %.2f\n", promedio)
}
