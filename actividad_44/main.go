/* Cargar una matriz A de 4x5, generar una función que devuelva como resultado un vector de 4 elementos,
donde cada elemento de sea la suma de una fila de A.
Generar otra función que devuelva un vector de 5 elementos, donde cada elemento sea la suma de una
columna de A.
Ambos procedimientos deben recibir como parámetro la matriz A. Imprimir los vectores generados. */

package main

import "fmt"

const (
	filas    = 2
	columnas = 2
)

var ()

func sumaFilas(A [filas][columnas]int) [filas]int {

	var (
		sumarFilas  int
		vectorFilas [filas]int
	)

	for i := 0; i < filas; i++ {
		sumarFilas = 0
		for j := 0; j < columnas; j++ {
			sumarFilas += A[j][i]
		}
		vectorFilas[i] = sumarFilas
	}

	return vectorFilas
}

func sumaColumnas(A [filas][columnas]int) [columnas]int {

	var (
		sumarColumnas int
		vectorFilas   [columnas]int
	)

	for i := 0; i < filas; i++ {
		sumarColumnas = 0
		for j := 0; j < columnas; j++ {
			sumarColumnas += A[i][j]
		}
		vectorFilas[i] = sumarColumnas
	}

	return vectorFilas
}

func main() {

	var (
		A [filas][columnas]int
	)

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Println("Ingrese un N°: ")
			fmt.Scanln(&A[i][j])
		}
	}

	fila := sumaFilas(A)
	fmt.Print(fila)
	columna := sumaColumnas(A)
	fmt.Print(columna)

}
