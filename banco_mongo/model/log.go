package model

import "gopkg.in/mgo.v2/bson"

//RegistroLog armazena quando a pagina foi visitada
type RegistroLog struct {
	ID         bson.ObjectId
	DataVisita string `json:"dataVisita" bson:"dataVisita"`
	Quem       string `json:"quem,omitempty" bson:"quem,omitempty"`
}
