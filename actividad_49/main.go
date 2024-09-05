/*Una casa de artículos deportivos almacena en una matriz A de 100x4 los datos correspondientes a los 100
artículos que vende, donde la primera columna contiene el código del artículo, la segunda el código de la
marca del artículo, la tercera la cantidad existente y la cuarta el precio de venta. Se pide:
	a) Generar un procedimiento que reciba como parámetro la matriz A e imprimir el código de artículo y el
	precio correspondiente de aquellos cuya cantidad existente supere las cincuenta unidades.
	b) Generar una función que reciba como parámetro la matriz A y devuelva como resultado una matriz con el
	código del artículo y el código de marca de aquellos artículos cuyo precio sea menor o igual a $300.
	c) Generar una función que ordene en forma ascendente la matriz A de acuerdo al precio de venta. La
	función deberá recibir como parámetro por referencia la matriz a ordenar y el número de columna de la
	matriz por la que se desea ordenar.
	d) Generar una función que imprima la matriz recibida como parámetro. Utilizarla para imprimir la matriz A.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	filas    = 10
	columnas = 4
)

// a) Generar un procedimiento que reciba como parámetro la matriz A e imprimir el código de artículo y el
// 	precio correspondiente de aquellos cuya cantidad existente supere las cincuenta unidades.

func stock50(A [filas][columnas]int) {

	for i := 0; i < filas; i++ {
		if A[i][2] > 50 {
			fmt.Println(A[i][0], " ", A[i][3])
		}
	}
}

// b) Generar una función que reciba como parámetro la matriz A y devuelva como resultado una matriz con el
// 	código del artículo y el código de marca de aquellos artículos cuyo precio sea menor o igual a $300.

func precio300(A [filas][columnas]int) [][]int {

	var (
		B [][]int
	)

	for i := 0; i < filas; i++ {
		if A[i][3] <= 300 {
			B = append(B, []int{A[i][0], A[i][1]})
		}
	}

	return B
}

/*c) Generar una función que ordene en forma ascendente la matriz A de acuerdo al precio de venta. La
función deberá recibir como parámetro por referencia la matriz a ordenar y el número de columna de la
matriz por la que se desea ordenar.*/

func ordenarMatriz() {}

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
	var A [filas][columnas]int

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			if j == 0 {
				idArticulo := rand.Intn(9000) + 1000
				A[i][j] = idArticulo
			} else if j == 1 {
				marca := rand.Intn(900) + 100
				A[i][j] = marca
			} else if j == 2 {
				stock := rand.Intn(90) + 10
				A[i][j] = stock
			} else {
				price := rand.Intn(900) + 100
				A[i][j] = price
			}
		}
	}

	imprimirMatriz(A)
	fmt.Print("Los artículos con stock mayor a 50 son: \n")
	stock50(A)

	fmt.Print("Los productos con precio menor a 300 son: \n")
	fmt.Print(precio300(A))

}
