package model

import "database/sql"

//Local armazena os dados da localidade pelo seu codigo telefonico
//Para SQLX funcionar os campos na tag db tem de ser em maiusculas
type Local struct {
	Pais             string         `json:"pais" db:"COUNTRY"`
	Cidade           sql.NullString `json:"cidade" db:"CITY"`
	CodigoTelefonico int            `json:"cod_telefone" db:"TELCODE"`
}
