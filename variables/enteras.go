package variables

import "fmt"
// metodo publico cuando escribo con mayuscula
func MuestroEnteros() {

	// si no e define un valor su valor es 0 asi es go
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}