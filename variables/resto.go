package variables

import (
	"fmt"
	"strconv"
	"time"
	
)

// PUEDO LLAMAR CUALQUIER FUNCION QUE PERTENEZCA A ESTE DIRECTORIO O package

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {

	// declaro una variable string aunque no es necsario declararla oouedo asignarla directamente
	
	// digo que esa variable tiene el valor del parametro pasado en la funcion que recibe un entero y pasara a string
	// con un paquete pude convertirlo
	texto := strconv.Itoa(numero)

	return true, texto
}
