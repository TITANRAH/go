// siempre es asi package impor func

package main

// importo el paquete variables y puedo acceder a todo su contenido
import (
	"fmt"
	"proyecto_go/variables"
)

func main() {

	// esto es como destructurar lo que devuelve la funcion la funcion devuelve texto y estado tomo lo que devuelve

	texto, estado := variables.ConviertoaTexto(39999)

	fmt.Println(texto)
	fmt.Println(estado)
}
