// Se tienen 2 vectores A y B de 20 elementos. Programar un procedimiento que permita ordenar
// vectores de forma ascendente. El procedimiento debe recibir como parámetro por referencia el
// vector a ordenar.
// Una vez ordenado los 2 vectores mediante la llamada al procedimiento, generar un nuevo vector C
// que sea la intercalación ordenada de A y de B (considere que no hay elementos repetidos).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cargaVector(vector *[20]int) {
	for i := 0; i < 20; i++ {
		vector[i] = rand.Intn(100)
	}
}

func ordenarVector(vector *[20]int) {

	n := len(vector)
	vaux := 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if vector[j] > vector[j+1] {
				vaux = vector[j+1]
				vector[j+1] = vector[j]
				vector[j] = vaux
			}
		}
	}
}

func intercalarVectores(A, B *[20]int) [40]int {

	va, vb := 0, 0
	var C [40]int

	for i := 0; i < 40; i++ {
		if va < 20 && (vb >= 20 || A[va] < B[vb]) {
			C[i] = A[va]
			va++
		} else {
			C[i] = B[vb]
			vb++
		}
	}
	return C
}

func main() {
	rand.Seed(time.Now().Unix())
	var A, B [20]int

	cargaVector(&A)
	ordenarVector(&A)
	fmt.Println(A)

	cargaVector(&B)
	ordenarVector(&B)
	fmt.Println(B)

	resultado := intercalarVectores(&A, &B)
	fmt.Println(resultado)
}
