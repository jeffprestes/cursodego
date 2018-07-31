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
func (i *Imovel) DobraValor() {
	i.valor = i.valor * 2
	return
}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//TriplicaValor triplica valor
func TriplicaValor(i *Imovel) {
	i.valor = i.valor * 3
}

func QuadruplicaValor(i Imovel) (x Imovel) {
	x.valor = i.valor * 4
	return
}
