package ejercicios

import (
	"strconv"
)

func PublicFunction(numeroString string) (string, int) {

	numeroConvertido, err := strconv.Atoi(numeroString)
	if err != nil {
		return "hubo un error" + err.Error(), 0
	}
	if numeroConvertido > 100 {

		return "Es mayor a 100", numeroConvertido
	} else {
		return "No es mayor a 100", numeroConvertido
	}

}
