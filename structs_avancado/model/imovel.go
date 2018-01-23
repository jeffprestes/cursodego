package model

//Imovel representa informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//DobraValor dobra o valor informado
func (i *Imovel) DobraValor(valor int) (resultado int) {
	i.valor = valor
	resultado = i.valor * 2
	return
}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
