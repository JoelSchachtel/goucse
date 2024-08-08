package main

import "fmt"

const (
	pi float64 = 3.1416
)

var (
	radio    float64
	longitud float64
	area     float64
)

func main() {
	fmt.Println("Ingrese el radio de la circunferencia:")
	fmt.Scanln(&radio)

	longitud = radio * 2 * pi
	area = pi * radio * radio
	fmt.Println("La longitud de la circunferencia es:", longitud)
	fmt.Println("El área de la circunferencia es:", area)
	fmt.Println("La longitud de la circunferencia es:", longitud, " y el área es", area)
}
