package main

import "fmt"

var (
	num int
	sum int
)

func main() {
	fmt.Print("Ingrese un número: ")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		sum += i

	}

	fmt.Printf("La suma de %d y sus anteriores es: %d", num, sum)

}
