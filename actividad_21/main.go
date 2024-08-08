/*Se ingresan N pares de números que indica el Número de móvil y la cantidad de cuadras recorridas de una
empresa de Remises. El precio por cuadra es de $1000 y la cantidad de autos son 6, existiendo una
correspondencia con el número de móvil. Se pide:
• Cantidad de viajes realizados por el móvil 3.
• Importe total recaudado por el móvil 5.
• Imprimir el número de móvil que realizó el viaje más largo y el importe recaudado.*/

package main

import "fmt"

const (
	valorDeCuadra int = 1000
)

var (
	movil             int
	cuadras           int
	viajes            int
	valorDeViaje      int
	viajeLargo        int
	importeViajeLargo int
	cuadrasViajeLargo int
	importe5          int
	viajes3           int
)

func main() {
	fmt.Println("Ingrese el número de móvil: ")
	fmt.Scan(&movil)

	for movil >= 1 && movil <= 6 {
		viajes = viajes + 1
		fmt.Println("Ingrese la cantidad de cuadras realizadas: ")
		fmt.Scan(&cuadras)

		valorDeViaje = cuadras * valorDeCuadra

		if movil == 3 {
			viajes3 = viajes3 + 1
		}

		if movil == 5 {
			importe5 = importe5 + valorDeViaje
		}

		if viajes == 1 {
			cuadrasViajeLargo = cuadras
			viajeLargo = movil
			importeViajeLargo = valorDeViaje
		} else if cuadrasViajeLargo < cuadras {
			viajeLargo = movil
			importeViajeLargo = valorDeViaje
		}

		fmt.Println("Ingrese el número de móvil: ")
		fmt.Scan(&movil)
	}

	fmt.Printf("El móvil 3 realizó %d viajes\n", viajes3)
	fmt.Printf("El móvil 5 reacaudó $%d\n", importe5)
	fmt.Printf("El móvil %d realizó el viaje mas largo y recaudó %d\n", viajeLargo, importeViajeLargo)

}
