// siempre es asi package impor func

package main

// importo el paquete variables y puedo acceder a todo su contenido
// import "proyecto_go/mapas"
// import "proyecto_go/users"
// import e "proyecto_go/ejer_interfaces"
// import "proyecto_go/modelos"
import "proyecto_go/defer_panic"

// "proyecto_go/files"
// "proyecto_go/funciones"
// "proyecto_go/arreglos_slice"
// "proyecto_go/mapas"
// "proyecto_go/ejercicios"
// "fmt"
// "proyecto_go/iteraciones"
// "proyecto_go/teclado"
// 	"proyecto_go/variables"
// 	"runtime"

func main() {

	/* testo, numero := ejercicios.PublicFunction("101")

	fmt.Println(testo)
	fmt.Println(numero)


	// esto es como destructurar lo que devuelve la funcion la funcion devuelve texto y estado tomo lo que devuelve

	texto, estado := variables.ConviertoaTexto(39999)

	fmt.Println(texto)
	fmt.Println(estado)

	// obtengo sistema operativo para sber si es linux window etc

	if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("esto no es windows", os)
	} else {
		fmt.Println("esto es windows", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")

	default:
		fmt.Println(os)
	}*/

	// num, text := ejercicios.PublicFunction("500");

	// fmt.Println(num)
	// fmt.Println(text);

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.Multiplicar())

	// files.SumTable()
	// funciones.LlamarClosure()

	// funciones.Exponencia(2)
	// users.AltaUsuario()

	// Pedro:=new(modelos.Hombre)
	// Roberto:=new(modelos.Hombre)
	// Roberto.Pensar()

	// e.HumanosRespirando(Roberto)

	// Maria:=new(modelos.Mujer)

	// e.HumanosRespirando(Maria)

	defer_panic.EjemploPanic()


}
