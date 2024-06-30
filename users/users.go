package users

import (
	"fmt"
	"proyecto_go/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
