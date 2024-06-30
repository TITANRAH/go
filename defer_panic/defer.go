package defer_panic

import (
	"fmt"
	"log"
)

func VamosDefer() {
	fmt.Println("Este es un primer mensaje")

	// defer es para una instruccion final osea lo ultimo que se ejecuta
	defer fmt.Println("Ese es el nsaje final")

	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {

	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encintro el valor 1")
	}
}
