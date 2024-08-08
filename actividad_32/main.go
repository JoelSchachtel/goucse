// Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que
// deberá pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones
// vendidas por cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad
// vendida y según esa categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

// Categoría Excursiones $ x Excursión
// A <10 $1000
// B >= 10 y < 100 $1500
// C >= 100 $1900

// Se ingresan las cantidades de excursiones vendidas de 10 promotores. Se necesita conocer:
// a) la cantidad de promotores de cada categoría
// b) el importe a pagarle a cada promotor
// Resolver:
//  Para el punto a) crear una función “categoriaPromotor” que recibirá como parámetro la
// cantidad de excursiones que vendió y devuelva como resultado la letra de la categoría a la que
// pertenece.
//  Para el punto b) crear una función “importeAPagar” que recibirá como parámetro la cantidad
// de excursiones que vendió, esta función deberá hacer uso de la función “categoriaPromotor”
// para determinar el importe a pagar.
// Imprimir los resultados.

package main

import "fmt"

var (
	viajesVendidos int
	categoria      string
	categoriaDef   string
	totalAPagar    int
)

func cargaPromotores() {
	fmt.Println("Ingrese la cantidad de excursiones vendidas:")
	fmt.Scanln(&viajesVendidos)
}

func categoriaPromotor(viajesVendidos int) string {

	if viajesVendidos < 10 {
		categoria = "A"
	} else if viajesVendidos > 10 && viajesVendidos < 100 {
		categoria = "B"
	} else {
		categoria = "C"
	}

	return categoria

}

func importeAPagar(viajesVendidos int) int {

	categoriaDef = categoria

	if categoriaDef == "A" {
		totalAPagar = viajesVendidos * 1000
	} else if categoriaDef == "B" {
		totalAPagar = viajesVendidos * 1500
	} else {
		totalAPagar = viajesVendidos * 1900
	}

	return totalAPagar
}

func main() {

	for i := 0; i < 10; i++ {
		cargaPromotores()

		categoria = categoriaPromotor(viajesVendidos)
		totalAPagar = importeAPagar(viajesVendidos)

		fmt.Printf("La categoría del promotor es %s\n", categoria)
		fmt.Printf("El importe a pagar al promotor es %d\n", totalAPagar)
	}

}
