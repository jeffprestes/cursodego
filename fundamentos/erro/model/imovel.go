package model

import "errors"

//Imovel representa um imovel
type Imovel struct {
	Nome  string `json:"nome"`
	valor int
	X     int `json:"coordenada_x"`
	Y     int `json:"coordenada_y"`
}

/*
ErrValorDeImovelInvalido é um erro que é apresentado quando é atribuido a imóvel um valor muito baixo
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para um imovel")

/*
ErrValorDeImovelMuitoAlto é um erro que é apresentado quando é atribuído ao imóvel um valor muito alto fora do mercado
*/
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto.")

//GetValor obtem o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor define o valor do Imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}
