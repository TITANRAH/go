package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

// no recibe ni devuelve es un metodos
func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1 :")
	if scanner.Scan() {

		// el : se utiliza cuando la variable no esta creada previamente
		// pero si las creo no es necesario el : , en estte caso las cree arriba
		numero1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			// esto devuelve el error es como throw o crear un error
			panic("Dato ingresado esta incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 :")

	if scanner.Scan() {

		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Dato ingresado esta incorrecto" + err.Error())
		}
	}
	fmt.Println("Ingrese leyenda 2 :")

	if scanner.Scan() {

		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)
}
