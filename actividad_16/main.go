package main

import "fmt"

var (
	nombre          string
	edad            int
	sueldo          int
	may25sueldo3000 string
	cont18          int
	sumasueldo      int
	empleados       int
)

func main() {
	fmt.Println("Ingrese la edad del empleado: ")
	fmt.Scanln(&edad)

	for i := 0; edad != 0; i++ {
		fmt.Println("Ingrese el nombre del empleado:")
		fmt.Scanln(&nombre)
		fmt.Println("Ingrese el sueldo del empleado")
		fmt.Scanln(&sueldo)
		empleados = i + 1
		sumasueldo = sumasueldo + sueldo
		if edad > 25 && sueldo > 3000 {
			may25sueldo3000 = may25sueldo3000 + nombre + " "
		} else if edad < 18 {
			cont18 += 1
		}

		fmt.Println("Ingrese la edad del empleado: ")
		fmt.Scanln(&edad)
	}

	fmt.Printf("Los empleados que tienen más de 25 y cobran más de 3000 son %s\n", may25sueldo3000)
	fmt.Printf("La cantidad de empleados menor a 18 es %d\n", cont18)
	fmt.Printf("La cantidad de empleados ingresados es: %d\n", empleados)
	fmt.Printf("La suma de los sueldos es: %d", sumasueldo)
}
