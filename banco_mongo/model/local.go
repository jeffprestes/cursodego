package model

//Local armazena os dados da localidade pelo seu codigo telefonico
type Local struct {
	Pais             string `json:"pais" db:"country" bson:"country"`
	Cidade           string `json:"cidade" db:"city" bson:"city"`
	CodigoTelefonico int    `json:"cod_telefone" db:"telcode" bson:"telcode"`
}
