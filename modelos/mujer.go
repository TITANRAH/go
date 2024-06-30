package modelos

type Mujer struct {
	// herencia de hombre
	Hombre
}

func (h *Mujer) Respirar() {
	h.repirando = true
}
func (h *Mujer) Comer() {
	h.comiendo = true
}
func (h *Mujer) Pensando() {
	h.pensando = true
}
func (h *Mujer) Sexo() string {
	return "Mujer"
}
