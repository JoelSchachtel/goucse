/*Teniendo como dato el número de matrícula y las notas de los 3 parciales evaluados, determinar si cada
alumno ingresado logró la condición de regularidad teniendo en cuenta:
• Las notas van del 1 al 10 y un parcial se aprueba con 7.
• Se regulariza la materia con al menos dos parciales aprobados.
• Un número de matrícula igual a 0 nos avisa el fin de datos.
Determinar además la cantidad total de alumnos, cuántos de ellos regularizaron la materia y el porcentaje
que éste grupo representa sobre el total. Imprimir los resultados.*/

package main

import "fmt"

var (
	matricula  int
	notas      [3]int
	c          int
	aprTotal   int
	porcentaje float32
	i          int
)

func main() {
	fmt.Println("Ingrese el número de matrícula: ")
	fmt.Scanln(&matricula)

	for i = 0; matricula != 0; i++ {
		apr := 0
		for c = 0; c < 3; c++ {
			fmt.Printf("Ingrese la nota %d: ", c+1)
			fmt.Scanln(&notas[c])
			if notas[c] <= 0 || notas[c] > 10 {
				fmt.Println("La nota ingresada no se encuentra en el rango 1 - 10")
				break
			}
			if notas[c] >= 7 {
				apr += 1
			}
		}
		if apr >= 2 {
			aprTotal += 1
			fmt.Println("El alumno esta regularizado.")
		} else {
			fmt.Println("Usted no está regularizado.")
		}
		fmt.Println("Ingrese el número de matrícula: ")
		fmt.Scanln(&matricula)
	}
	fmt.Println("La cantidad de alumnos ingresados fue: ", i)
	porcentaje = (float32(aprTotal) * 100 / float32(i))
	fmt.Printf("\n El porcentaje de alumnos regularizados es: %.2f%%", porcentaje)
}
