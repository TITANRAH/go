package funciones

import "fmt"
// funciones asignadas a variables o enviadas por parametros
func Calculos() {
	// var numero1, numero2 int
	var numero3 int = 32 
	var numero4 int = 243

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 
	}

	fmt.Println(calculo(10, 25))
}