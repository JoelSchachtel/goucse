/*Se lee la siguiente información: Sexo, Estado Civil y Edad en donde: sexo podrá tomar los valores H o M,
estado civil S (soltero), E (separado), C (casado) y V (viudo). Se pide:
• Calcular el promedio de edad de los hombres casados. X
• Contar la cantidad de hombres casados mayores de 28 años. X
• Contar la cantidad de hombres viudos X
• Contar la cantidad de mujeres solteras menores de 25 años. X
• Contar cuantos hombres y mujeres se ingresaron. X
El final de la lectura viene indicado por un “S” en el campo sexo.*/

package main

import "fmt"

var (
	sexo            string
	edad            int
	eCivil          string
	hcm28           int
	hombres         int
	edadHombres     int
	hombresCasados  int
	hombresViudos   int
	mujeres         int
	mcm25           int
	promedioHombres int
)

func main() {
	fmt.Println("Ingrese el sexo:")
	fmt.Scan(&sexo)

	for sexo != "S" {
		fmt.Println("Ingrese la edad:")
		fmt.Scan(&edad)
		fmt.Println("Ingrese el estado civil:")
		fmt.Scan(&eCivil)

		if sexo == "H" || sexo == "h" {
			hombres += 1
			if eCivil == "C" || eCivil == "c" {
				hombresCasados += 1
				edadHombres = edad + edadHombres
				promedioHombres = edadHombres / hombres
				if edad > 28 {
					hcm28 += 1
				} else if eCivil == "V" || eCivil == "v" {
					hombresViudos += 1
				}
			}
		} else if sexo == "M" || sexo == "m" {
			mujeres += 1
			if eCivil == "S" {
				mcm25 += 1
			}
		}
		fmt.Println("Ingrese el sexo:")
		fmt.Scan(&sexo)
	}

	fmt.Printf("La cantidad de hombres es: %d y la cantidad de mujeres es: %d\n", hombres, mujeres)
	fmt.Printf("La cantidad de hombres casados es: %d\n", hombresCasados)
	fmt.Printf("La cantidad de hombres casados mayores a 28 es: %d\n", hcm28)
	fmt.Printf("El promedio de la edad de hombres ingresados es: %d\n", promedioHombres)
	fmt.Printf("La cantidad de hombres viudos es: %d\n", hombresViudos)
	fmt.Printf("La cantidad de mujeres solteras mayores a 25 es: %d\n", mcm25)

}
