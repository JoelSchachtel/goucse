/*
En una empresa se utiliza una matriz A de 51x6 elementos para guardar información correspondiente a las
ventas de 50 vendedores en los 4 trimestres del año. La primera columna guarda información correspondiente
al número de vendedor, a partir de la segunda columna y hasta la quinta se guardan las ventas de ese
vendedor para cada trimestre. Se pide:
	a) Calcular e imprimir el total de ventas
	• de cada vendedor a lo largo del año
	• de cada trimestre
	• anual [X]
	b) Ordenar la matriz A en forma descendente de acuerdo al total de ventas de cada vendedor.
	c) Ingresar un número de vendedor e imprimir en qué trimestre realizó la mayor y la menor venta.
Nota: La fila 51 y la columna 6 pueden ser utilizadas para los fines que crea conveniente.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filas    = 51
	columnas = 6
)

func imprimirMatriz(A [filas][columnas]int) {

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Print(A[i][j], " ")
		}
		fmt.Println("")
	}
}

func main() {

	rand.Seed(time.Now().Unix())
	var (
		A [filas][columnas]int
	)

	for i := 0; i < filas; i++ {
		A[i][0] = i + 1
		for j := 1; j <= 4; j++ {
			A[i][j] = rand.Intn(100) + 1
		}
		A[i][5] = A[i][1] + A[i][2] + A[i][3] + A[i][4]
	}

	fmt.Println("Planilla inicial de ventas: ")
	imprimirMatriz(A)
}
