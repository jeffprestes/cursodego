package main

import (
	"fmt"

	"github.com/jeffprestes/cursodego/interfaces/model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmaPataNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
