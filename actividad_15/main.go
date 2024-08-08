package main

import "fmt"

var (
	nombre      [10]string
	edad        [10]int
	nombremayor string
	nombremenor string
	edadmayor   int
	edadmenor   int
	prom        int
	cont12      int
	cont1016    int
)

func main() {
	fmt.Println("Ingrese el nombre y la edad de 10 alumnos.")

	for i := 0; i < 3; i++ {
		fmt.Printf("Nombre %d: ", i+1)
		fmt.Scanln(&nombre[i])
		fmt.Printf("Edad %d: ", i+1)
		fmt.Scanln(&edad[i])
		prom = prom + edad[i]

		if i == 0 {
			nombremayor = nombre[i]
			edadmayor = edad[i]
			nombremenor = nombre[i]
			edadmenor = edad[i]
		} else if edad[i] > edadmayor {
			nombremayor = nombre[i]
			edadmayor = edad[i]
		} else if edad[i] < edadmenor {
			nombremenor = nombre[i]
			edadmenor = edad[i]
		}

		if edad[i] > 12 {
			cont12 += 1
		}

		if edad[i] > 10 && edad[i] < 16 {
			cont1016 += 1
		}
	}

	prom = prom / 3

	fmt.Printf("El de mayor edad es %s, y el de menor edad es %s\n", nombremayor, nombremenor)
	fmt.Printf("La cantidad de alumnos mayores a 12 es %d, y entre 10 y 16 es %d\n", cont12, cont1016)
	fmt.Printf("El promedio de las edades es %d", prom)
}
