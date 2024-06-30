package arreglos_slice

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 123, 2, 4, 5, 6, 7, 8, 2, 1}

func MuestroSlice() {

	// fmt.Println(tablaS)

	// si no indicamos una posicon a la izquiedda  derecha sera desde que comienza o hasta el final
	// segun la omision sea izquierda o derecha

	porcion := arreglo[3:]   //slice creado desde un vector dede la posicion 3
	porcion2 := arreglo[:5]  //slice creado desde el prinpicio hasta la 5
	porcion3 := arreglo[6:7] //slice cread	o desde la posicion 6 a 7

	fmt.Println(arreglo)
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {

	// esto creara un array de 5 pero con capacidad 20
	elementos := make([]int, 5, 20)

	// defino d y le paso elementos a cada d
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	// puedo definir asi y go dara la capacidad
	nums := make([]int, 0)

	for i := 0; i < 100; i++ {

		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
