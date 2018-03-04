package model

//Usuario representa um usuario do sistema
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
