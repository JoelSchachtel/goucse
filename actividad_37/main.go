// Cargar una matriz A de 6x5, Crear una función que reciba como parámetro la matriz cargada y devuelva la suma
// de aquellos elementos que sean mayor o igual que 5 y la cantidad de elementos que intervinieron. Imprimir los
// resultados.

package main

import "fmt"

var (
	secuencia     [6][5]int
	numeroSumados int
	contadorMayor int
)

func sumaMatriz(matriz [6][5]int) (int, int) {

	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			if matriz[i][j] >= 5 {
				numeroSumados = numeroSumados + matriz[i][j]
				contadorMayor = contadorMayor + 1
			}
		}
	}
	return contadorMayor, numeroSumados
}

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("Ingrese un N°")
			fmt.Scanln(&secuencia[i][j])
		}
	}
	sumaMatriz(secuencia)
	fmt.Printf("Los números mayores de 5 son %d y su suma da %d\n", contadorMayor, numeroSumados)
	fmt.Println(secuencia)
}
