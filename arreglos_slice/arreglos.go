package arreglos_slice

import "fmt"

// crea un arreglo de 10 eleemntos enteros
// asignacion añ crear el arreglo, los primeros 3 puestos
var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {

	tabla[2] = 54
	tabla[7] = 33

	// tabla2 := [10]string{"Pablo", "Juan"}
	// fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		// fmt.Println((tabla[i]))
	}

	matriz[7][24] = 15


	fmt.Println(matriz)
}
