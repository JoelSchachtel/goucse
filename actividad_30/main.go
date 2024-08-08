package main

import "fmt"

var (
	dolaresAComprar float32
	precioCoti      float32
	pesosArgDisp    float32
	puedo           bool
	consulta        float32
)

func cargaPesos() {
	fmt.Println("Ingrese la cantidad de pesos disponibles:")
	fmt.Scan(&pesosArgDisp)

}

func dolAComprar() {
	fmt.Println("Ingrese la cantidad de dolares a comprar:")
	fmt.Scan(&dolaresAComprar)
}

func cotizacion() { // En base a la cantidad de dolares y de pesos realiza la cotizacion
	fmt.Println("Ingrese la cotización del dia:")
	fmt.Scan(&precioCoti)
}

func cotizador(dolAComprar, pesosArgDisp, precioCoti float32) bool {

	consulta = dolAComprar * precioCoti

	if consulta < pesosArgDisp {
		return true
	} else {
		return false
	}

}

func main() { //Funcion principal del programa
	cargaPesos()
	dolAComprar()
	cotizacion()

	for pesosArgDisp != 0 {
		puedo = cotizador(dolaresAComprar, pesosArgDisp, precioCoti)

		if puedo {
			fmt.Println("Puedo comprar dolares")
		} else {
			fmt.Println("No puedo comprar dolares")
		}
		cargaPesos()
	}

}
