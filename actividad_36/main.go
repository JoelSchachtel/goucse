// Se pide:
//  Leer los datos de los viajes. El fin de carga viene dado por un número de móvil igual a cero. (LISTO)
//  Cargar un vector REM de 6 elementos con el total de cuadras que hizo cada móvil. (LISTO)
//  Crear una función que reciba como parámetro un número de móvil y el vector RES e imprima
// la cantidad total de cuadras realizadas y el importe recaudado por ese móvil.
//  Utilizar la función del punto anterior para imprimir en pantalla el siguiente reporte:
// Nro. móvil Cant. total cuadras Importe recaudado
// 1 100 8000
// 2 150 12000
// 3 90 7200
// 4 200 16000
// 5 205 16400
// 6 90 7200
//  Realizar un procedimiento que permita ordenar un vector de forma descendente. El
// procedimiento deberá recibir como parámetro por referencia el vector a ordenar.
//  Ordenar el vector REM con el procedimiento desarrollado en el punto anterior e imprimir el
// vector ordenado.

package main

import "fmt"

var (
	cuadras int
	nMovil  int
	REM     [6]int
)

func calculadoraDeCuadras(movil int, vector *[6]int) {
	var cantCuadras, recaudado int

	cantCuadras = vector[movil-1]
	recaudado = cantCuadras * 80

	fmt.Printf("%d - %d - %d\n", movil, cantCuadras, recaudado)
}

func ordenarVector(vector *[6]int) {
	vaux := 0
	for i := 0; i < 6-1; i++ {
		for j := 0; j < 6-i-1; j++ {
			if vector[j] > vector[j+1] {
				vaux = vector[j+1]
				vector[j+1] = vector[j]
				vector[j] = vaux
			}
		}
	}
}

func main() {
	fmt.Println("Ingrese el N° de Móvil")
	fmt.Scanln(&nMovil)

	for nMovil != 0 {
		if nMovil < 0 || nMovil > 6 {
			fmt.Println("Ingrese un móvil válido. 1 - 6")

		} else {
			fmt.Println("Ingrese la cantidad de cuadras realizadas")
			fmt.Scanln(&cuadras)
			REM[nMovil-1] = REM[nMovil-1] + cuadras

		}

		fmt.Println("Ingrese el N° de Móvil")
		fmt.Scanln(&nMovil)
	}

	fmt.Println(REM)

	for x := 0; x < 6; x++ {
		calculadoraDeCuadras(x+1, &REM)
	}

	ordenarVector(&REM)
	fmt.Println(REM)

}
