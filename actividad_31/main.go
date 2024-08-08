// Leer 5 números e indicar por cada uno si es múltiplos de 3 y no de 5. Para ello crear una función que se
// llame “verificarMultiplo” que reciba 2 parámetro:
//	 	* El número que se desea verificar si es múltiplo de.
//	 	* El múltiplo a verificar (3 o 5).
// La función deberá devolver como resultado un valor booleano indicando si es múltiplo o no. Para
// verificar si la multiplicidad se cumple se deberá realizar la división restando.
// El programa a desarrollar deberá implementar la función para resolver el problema planteado.

package main

import "fmt"

var (
	numero int
	resto1 bool
	resto2 bool
)

func leer() {

	fmt.Println("Ingrese un N°:")
	fmt.Scanln(&numero)

}

func verificarMultiplo(numero int, multiploaverf int) bool {

	if numero%multiploaverf == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	for i := 0; i < 5; i++ {
		leer()

		resto1 = verificarMultiplo(numero, 3)
		resto2 = verificarMultiplo(numero, 5)

		if resto1 && !resto2 {
			fmt.Printf("%d es múltiplo de 3 y NO de 5.\n", numero)
		} else if resto1 && resto2 {
			fmt.Printf("%d es múltiplo de 3 y de 5.\n", numero)
		} else if !resto1 && resto2 {
			fmt.Printf("%d NO es múltiplo de 3, pero si de 5.\n", numero)
		} else {
			fmt.Print("No es múltiplo de 3 ni de 5.\n")
		}
	}
}
