package main

import "fmt"

const (
	iva float64 = 0.21
)

func main() {
	var (
		cantidad int
		punit    float64
		subtotal float64
		piva     float64
		factura  float64
	)
	fmt.Println("Ingrese la cantidad del producto:")
	fmt.Scanln(&cantidad)
	fmt.Println("Ingrese el precio unitario del producto:")
	fmt.Scanln(&punit)
	subtotal = float64(cantidad) * punit
	fmt.Println("El subtotal es $", subtotal)
	piva = subtotal * iva
	fmt.Println("El IVA es: $", piva)
	factura = subtotal + piva
	fmt.Println("El total de la factura es: $", factura)
}
