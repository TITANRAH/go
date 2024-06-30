package funciones

import "fmt"

// funcion table devuelve una funcion que a su vez dvuelve un entero
func table(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	MiTabla := table(tabladel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
