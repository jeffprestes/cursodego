package model

type movel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	Valor int
}

//NovoMovel encapsula a criacao da instancia de novo movel
func NovoMovel(x, y, valor int, nome string) movel {
	m := movel{x, y, nome, valor}
	return m
}
