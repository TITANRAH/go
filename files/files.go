package files

import (
	// "bufio"
	"bufio"
	"fmt"
	// "io/ioutil"
	// "strings"

	// "ioutil"
	"os"
	"proyecto_go/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func SaveTable() {
	// llamo a la funcion que pide un numero y genera una tabla
	var texto string = ejercicios.Multiplicar()
	// creo un buffer un archivo para guardar txt
	archivo, err := os.Create(filename)
	// si da error retorna
	if err != nil {
		fmt.Println("Error al crear el arcchivo" + err.Error())
		return
	}
	// si no da error grabara con Fprintln
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumTable() {

	// pide el numero
	var texto string = ejercicios.Multiplicar()

	// realiza el guardado pide filename declarado arriba y el texto generado por tabla multipilicar
	if !Append(filename, texto) {
		fmt.Println("Error al concatenear el contenido ")
	}

}

func Append(filen string, texto string) bool {

	// OpenFile permite abrir el archivo y asi concatenar nuevo contenido
	// PUEDO SEPARAR ARGUMENTOS CON UN PIPE SIENDO QUE LA FUNCION RECIBE SOLO 2 ARGYUMENTOS
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	// esta funcion devuelve dos parametros, omito el primero y tomo el error
	// realiza el guardado
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	// cierro el archivo
	arch.Close()
	return true
}

func LeoArchivo() {

	// archivo, err := os.ReadFile(filename)
	archivo, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
	}

	// fmt.Println(string(archivo))
	// fmt.Println(archivo)
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		register := scanner.Text()
		fmt.Println(">" + register)
	}
}


// 1. crear el codigo que pide un numero y genra una tabla de multiplicar 
// 2. llamar a esa funcion aca y usar save tabla que pedira un numero y guardara la data en el archivo asignado
// 3. crear funcion apepend que permite sumar data en vez de reempalzar el archivo
// 4. llamar a append dentro de sumtabley ahora usaremos sumtable