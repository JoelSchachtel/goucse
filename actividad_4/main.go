package main

import "fmt"

func main() {
	var dia int8

	fmt.Println("Ingrese un número para determinar el día de la semana (Del 1 al 7):")
	fmt.Scanln(&dia)

	switch dia {
	case 0:
		fmt.Println("El número ingresado", dia, "no corresponde a ningún dia")
	case 1:
		fmt.Println("El día ingresado es: Lunes")
	case 2:
		fmt.Println("El día ingresado es: Martes")
	case 3:
		fmt.Println("El día ingresado es: Miércoles")
	case 4:
		fmt.Println("El día ingresado es: Jueves")
	case 5:
		fmt.Println("El día ingresado es: Viernes")
	case 6:
		fmt.Println("El día ingresado es: Sábado")
	case 7:
		fmt.Println("El día ingresado es: Domingo")
	case 8:
		fmt.Println("El número ingresado", dia, "no corresponde a ningún dia")
	case 9:
		fmt.Println("El número ingresado", dia, "no corresponde a ningún dia")
	}
}
