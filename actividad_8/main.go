package main

import "fmt"

func main() {
	var (
		hora     string
		minutos  string
		segundos string
	)

	fmt.Println("Ingrese la hora:")
	fmt.Scanln(&hora)
	fmt.Println("Ingrese los minutos:")
	fmt.Scanln(&minutos)
	fmt.Println("Ingrese los segundos:")
	fmt.Scanln(&segundos)

	fmt.Println(hora, ":", minutos, ":", segundos)
}
