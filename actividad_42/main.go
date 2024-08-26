/*
Cargar una matriz A de 10x20 elementos. Generar una función que devuelva como resultado el mayor
elemento de las filas pares y el menor de las filas impares y las posiciones donde se ubican. La función deberá
recibir como parámetros:
	a) La matriz a buscar
	b) Si se desea buscar el mayor o menor (parámetro booleano)
	c) Si se desea buscar sobre filas pares o impares (parámetro booleano)
Imprimir los resultados.
*/

package main

import (
	"fmt"
)

const (
	fila    = 2
	columna = 4
)

func buscador(A [fila][columna]int, mayor bool, par bool) (valorReferencia, posicioni, posicionj int) {

	inicio := 0

	if !par {
		inicio = 1

	}

	for i := inicio; i < fila; i += 2 {
		for j := 0; j < columna; j++ {
			if i == inicio && j == 0 {
				valorReferencia = A[i][j]
				posicioni = i
				posicionj = j
			} else if mayor && A[i][j] > valorReferencia || !mayor && A[i][j] < valorReferencia {
				valorReferencia = A[i][j]
				posicioni = i
				posicionj = j
			}

		}
	}

	return valorReferencia, posicioni, posicionj
}

func main() {

	var A [fila][columna]int

	for i := 0; i < fila; i++ {
		for j := 0; j < columna; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&A[i][j])
		}
	}

	mayorPar, posiParI, posiParJ := buscador(A, true, true)
	menorImpar, posiImparI, posiImparJ := buscador(A, false, false)

	fmt.Printf("El mayor valor en filas pares es %d en la posición [%d][%d]\n", mayorPar, posiParI, posiParJ)
	fmt.Printf("El menor valor en filas impares es %d en la posición [%d][%d]\n", menorImpar, posiImparI, posiImparJ)
}
