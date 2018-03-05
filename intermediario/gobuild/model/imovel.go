package model

//Imovel representa um imovel
type Imovel struct {
	Nome  string `json:"nome"`
	valor int
	X     int `json:"coordenada_x"`
	Y     int `json:"coordenada_y"`
}

//GetValor obtem o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor define o valor do Imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}
