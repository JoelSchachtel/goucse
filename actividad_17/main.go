/* En una elección existen 3 candidatos. Por cada elector se tienen los siguientes datos: sexo (H o M) y el
partido por el cual votó (1 a 3). Ingresar “n” electores, determinar e imprimir:
• La cantidad de mujeres que votaron.
• La cantidad de varones que votaron al candidato 3.
• Candidato que ganó la elección, considerando que no hay empate. */

package main

import "fmt"

var (
	sexo        string
	voto        int
	cantmujeres int
	vot1        int
	vot2        int
	vot3        int
	h3          int
	electores   int
)

func main() {

	fmt.Println("Ingrese la cantidad de electores:")
	fmt.Scanln(&electores)

	for i := 0; i < electores; i++ {

		fmt.Println("Ingrese sexo: H para Hombre y M para mujer.")
		fmt.Scanln(&sexo)
		fmt.Println("Ingrese el voto: 1 - 2 - 3")
		fmt.Scanln(&voto)

		if sexo == "m" || sexo == "M" {
			cantmujeres += 1
		} else if sexo == "H" || sexo == "h" && voto == 3 {
			h3 += 1
		}

		switch voto {
		case 1:
			vot1 += 1
		case 2:
			vot2 += 1
		case 3:
			vot3 += 1
		}

	}

	if vot1 > vot2 && vot1 > vot3 {
		println("El ganador es el candidato 1 con: ", vot1, " votos")
	} else if vot2 > vot3 {
		println("El ganador es el candidato 2, con: ", vot2, " votos")
	} else {
		println("El ganador es el candidato 3, con:", vot3, " votos")
	}

	fmt.Println("La cantidad de votantes mujeres es:", cantmujeres)
	fmt.Println("Votantes hombres que votaron al candidato 3: ", h3)

}
