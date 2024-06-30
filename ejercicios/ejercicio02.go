package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// no hace falta ponerparentesis de la devolucion cuando es 1 parmetro
func Multiplicar() string {
	var texto string
	var numero1 int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número: ")

		if scanner.Scan() {

			numero1, err = strconv.Atoi(scanner.Text())
			// nil es nulo
			if err != nil {
				fmt.Println("Puedes ingresar solo números vuelve a intentarlo")
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		result := numero1 * i

		// Sprintf devuelve un string
		// text = text + fmt...
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, result)
	}

	

	return texto

}
