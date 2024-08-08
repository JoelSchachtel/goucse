// Realizar un programa que permita jugar un número indeterminado de veces a “Adivina el Número”
// que consiste en lo siguiente:
// 	* Genere un número aleatoriamente comprendido entre 0-100.
//  * Permitir al jugador adivinar el número generado, para ello deberá introducir por teclado
// 	números mientras no acierte. Para ello desarrollar una función que valide si el usuario adivinó
//	 el número o no, la función deberá recibir como parámetro el número ingresado por el usuario y
// 	el número generado aleatoriamente e informar por pantalla si el número generado es mayor o
// 	menor que el ingresado. Además deberá devolver un valor booleando que indique si acertó el
// 	número o no.
//  * Contar el número de intentos.
//  * En el momento de la victoria, mostrarlo en pantalla e informar el número de intentos que ha
// utilizado, sí el número de intentos es:
// o <= 5 se indica al jugador que es bueno.
// o 5 < número de intentos < 15 se indica al jugador que es regular.
// o el número de intentos es > 15 se indica que el jugador no es muy bueno.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	numeroUsuario, numeroParaAdivinar, intentos int
	jugador                                     string
	adivina                                     bool
)

func generarNumero() {
	numeroParaAdivinar = rand.Intn(100)
}

func adivinarNumero(numeroUsuario, numeroParaAdivinar int) bool {
	if numeroUsuario > numeroParaAdivinar {
		println("Ingrese un número menor.")
		adivina = false

	} else if numeroUsuario < numeroParaAdivinar {
		println("Ingrese un número mayor.")
		adivina = false
	} else {
		adivina = true
	}

	return adivina
}

func tipoDeJugador(intentos int) string {
	if intentos <= 5 {
		jugador = "BUEN"
	} else if intentos > 5 && intentos < 15 {
		jugador = "REGULAR"
	} else {
		jugador = "NO MUY BUEN"
	}

	return jugador
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Bienvenido al juego: ADIVINA EL NÚMERO.")
	generarNumero()
	// println(numeroParaAdivinar)
	fmt.Println("Número aleatorio generado. Ingresa el primer número:")
	fmt.Scanln(&numeroUsuario)

	for !adivinarNumero(numeroUsuario, numeroParaAdivinar) {
		fmt.Println("Ingrese otro número:")
		fmt.Scanln(&numeroUsuario)
		intentos = intentos + 1
	}
	tipoDeJugador(intentos)

	if adivina {
		fmt.Printf("El número ingresado es el correcto. GANASTE. Realizaste %d intentos, eres un %s jugador", intentos, jugador)
	}

}
