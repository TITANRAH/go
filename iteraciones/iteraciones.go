package iteraciones

import (
	"fmt"
)

func Iterar() {
	
	// significaria que dara saltos de a 5
	// i+=5
	for i:=0;i<10;i++{

		if i == 6 {
			// deteiene y llega hasta el 6
			// break

			// continue se salta el 6
			continue

		}

		fmt.Println(i)
		
	}
}
