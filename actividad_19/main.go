/* Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que deberá
pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones vendidas por
cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad vendida y según esa
categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

Categoría Excursiones $ x Excursión
A <10 $10
B >= 10 y < 50 $15
C >= 50 y < 100 $17
D >= 100 $19

Se cuenta con N datos que corresponden a las cantidades vendidas por cada promotor. Se necesita
conocer:
• El importe a pagarle a cada promotor
• La cantidad de promotores de cada categoría */

package main

import "fmt"

var (
	vendedores int
	viajes     int
	importe    int
	cata       int
	catb       int
	catc       int
	catd       int
)

func main() {
	fmt.Println("Ingrese la cantidad de vendedores:")
	fmt.Scanln(&vendedores)

	for i := 0; i < vendedores; i++ {
		fmt.Printf("Ingrese la cantidad de viajes vendidos por el promotor %d: ", i+1)
		fmt.Scanln(&viajes)
		if viajes < 10 {
			importe = viajes * 10
			fmt.Printf("Al promotor %d hay que pagarle %d \n", i+1, importe)
			cata += 1

		} else if viajes < 50 {
			importe = viajes * 15
			fmt.Printf("Al promotor %d hay que pagarle %d \n", i+1, importe)
			catb += 1
		} else if viajes < 100 {
			importe = viajes * 17
			fmt.Printf("Al promotor %d hay que pagarle %d \n", i+1, importe)
			catc += 1

		} else if viajes > 100 {
			importe = viajes * 19
			fmt.Printf("Al promotor %d hay que pagarle %d \n", i+1, importe)
			catd += 1
		}
	}

	fmt.Printf("En la categoría A hay: %d \n", cata)
	fmt.Printf("En la categoría B hay: %d \n", catb)
	fmt.Printf("En la categoría C hay: %d \n", catc)
	fmt.Printf("En la categoría D hay: %d \n", catd)

}
