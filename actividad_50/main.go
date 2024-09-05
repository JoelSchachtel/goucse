/* 21) Cargar una matriz A de Nx4 elementos, donde cada fila contiene los datos correspondientes a un libro de una
librería (la primera columna contiene el código del libro, la segunda el código del autor, la tercera el número de
ejemplares y la cuarta el precio del libro). Se pide:

		a) Generar un procedimiento que reciba como parámetro un código de un autor e imprima el código de todos
		los libros que éste publicó. El código de autor que se debe pasar como parámetro debe ser leído antes de
		invocar al procedimiento. [X]

		b) Generar una función que devuelva el código del libro y el código del autor del libro más caro. Imprimir los
		resultados. [X]

		c) Generar una función que devuelva una matriz que contenga el código de libro y el código del autor de
		aquellos libros cuyos números de ejemplares sea mayor a 45 unidades. [X]

		d) Ordenar la nueva matriz generado por código de autor. [X]

		e) Imprimir la matriz generada. [X] */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	columnas = 4
)

func librosDelAutor(A [][columnas]int, codigoAutor int) []int {

	var autor []int

	for i := 0; i < len(A); i++ {
		if codigoAutor == A[i][1] {
			autor = append(autor, A[i][0])
		}
	}
	return autor
}

func imprimirMatriz(A [][]int) {

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			fmt.Print(A[i][j], " ")
		}
		fmt.Println("")
	}
}

func masCaro(A [][columnas]int) (int, int) {

	var (
		libroMasCaro = A[0][3]
		codLibro     = A[0][0]
		codAutor     = A[0][1]
	)

	for i := 1; i < len(A); i++ {
		if A[i][3] > libroMasCaro {
			libroMasCaro = A[i][3]
			codLibro = A[i][0]
			codAutor = A[i][1]
		}
	}
	return codLibro, codAutor
}

func stockMayor45(A [][columnas]int) [][]int {

	var (
		B [][]int
	)

	for i := 0; i < len(A); i++ {
		if A[i][2] >= 45 {
			B = append(B, []int{A[i][0], A[i][1]})
		}
	}

	return B

}

func ordenarMatriz(B [][]int) {

	for i := 0; i < len(B)-1; i++ {
		for j := 0; j < len(B)-1-i; j++ {
			if B[j][1] > B[j+1][1] {
				B[j], B[j+1] = B[j+1], B[j]
			}
		}
	}

}

func main() {

	var (
		libros      int
		A           [][columnas]int
		codigoAutor int
	)

	rand.Seed(time.Now().Unix())

	fmt.Println("Ingrese la cantidad de libros a cargar:")
	fmt.Scanln(&libros)

	for i := 0; i < libros; i++ {
		var filasM [columnas]int        //La cantidad de columnas, definen el largo del vector
		for j := 0; j < columnas; j++ { // Recorre las columnas para irlas cargando
			switch j {
			case 0:
				filasM[j] = rand.Intn(100) + 1
			case 1:
				filasM[j] = rand.Intn(100) + 1
			case 2:
				filasM[j] = rand.Intn(100) + 1
			case 3:
				filasM[j] = rand.Intn(100) + 1
			}
		}
		A = append(A, filasM) //Agrega el vector a la matriz
	}

	//a) Libros del autor
	fmt.Println("Ingrese el código del autor:")
	fmt.Scanln(&codigoAutor)
	fmt.Print(librosDelAutor(A, codigoAutor), "\n")

	//b) Mas caro
	codLibro, codAutor := masCaro(A)
	fmt.Printf("El libro mas caro es el código %d del autor %d \n", codLibro, codAutor)

	//c) Libros con stock mayor a 45
	fmt.Println("Los productos con stock mayor a 45 son:")
	mayorStock := stockMayor45(A)
	imprimirMatriz(mayorStock)

	//d) Ordenar la matriz
	ordenarMatriz(mayorStock)

	//e) Imprimir la matriz ordenada
	fmt.Println("Matriz ordenada por código de autor:")
	imprimirMatriz(mayorStock)
}
