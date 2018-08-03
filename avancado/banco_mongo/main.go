package main

import (
	"fmt"
	"net/http"

	"github.com/jeffprestes/cursodego/avancado/banco_mongo/manipulador"
	"github.com/jeffprestes/cursodego/avancado/banco_mongo/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	err = repo.AbreSessaoComMongo()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir a sessao com o MongoDB: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
