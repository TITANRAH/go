package funciones

import "fmt"

func Exponencia(valor int) {
	// al llamarla abajo de este if al llamar a esta funcion y pasarle un parametro 
	// hara la multiplicacion por 2 hasta que el valor sea mayor a 100000000 quiere decir que el valor geberado 
	// que despyes es el mismo valor pasado por paranetro no sea nmayor que 10000000
	if valor > 10 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}
