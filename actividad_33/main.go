// Realizar un programa que lea dos números A y B y que le pregunte al usuario que operación desea
// realizar. Las operaciones son: 1 = sumar, 2 = restar, 3 = multiplicar, 4 = dividir, 0 = salir.
// Una vez que el usuario haya ingresado la operación a realizar y los números a operar, el programa
// deberá mostrar el resultado de la operación y volver a preguntar al usuario que desea hacer,
// mientras que el usuario no ingrese la opción 0 = salir, el programa deberá seguir funcionando.
// Cada operación deberá ser implementada con distintas funciones y deberá recibir como parámetro
// los números a operar y devolver el resultado de la operación.

package main

import "fmt"

var (
	numero1, numero2, operacion, resultado float32
)

func suma(numero1, numero2 float32) float32 {
	resultado = numero1 + numero2
	return resultado
}

func resta(numero1, numero2 float32) float32 {
	resultado = numero1 - numero2
	return resultado
}

func division(numero1, numero2 float32) float32 {
	resultado = numero1 / numero2
	return resultado
}

func multiplicacion(numero1, numero2 float32) float32 {
	resultado = numero1 * numero2
	return resultado
}

func main() {
	fmt.Println("Ingrese el tipo de operación: (1- Suma // 2- Resta // 3- Multiplicación // 4- División // 0- Salir)")
	fmt.Scanln(&operacion)

	for operacion != 0 {

		fmt.Println("Ingrese el primer número:")
		fmt.Scanln(&numero1)

		fmt.Println("Ingrese el segundo número:")
		fmt.Scanln(&numero2)

		if operacion == 1 {
			resultado = suma(numero1, numero2)
		} else if operacion == 2 {
			resultado = resta(numero1, numero2)
		} else if operacion == 3 {
			resultado = multiplicacion(numero1, numero2)
		} else if operacion == 4 {
			resultado = division(numero1, numero2)
		} else {
			fmt.Println("El número ingresado no responde a ninguna operación.")
		}
		fmt.Printf("El resultado de la operación es: %f\n", resultado)
		fmt.Println("Ingrese el tipo de operación: (1- Suma // 2- Resta // 3- Multiplicación // 4- División // 0- Salir)")
		fmt.Scanln(&operacion)
	}

}
