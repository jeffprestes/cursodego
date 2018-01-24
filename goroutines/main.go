package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"sync"

	"github.com/jeffprestes/cursodego/goroutines/model"
)

var orquestrador sync.WaitGroup

func main() {
	orquestrador.Add(2)
	log.Println("****** Começando...")
	go traduzirParaJSON("saopaulo")
	go traduzirParaJSON("riodejaneiro")
	orquestrador.Wait()
	log.Println("****** Acabei! ******")
}

func traduzirParaJSON(nomeArquivo string) {
	fmt.Println(time.Now(), " - Começando a tradução do arquivo: ", nomeArquivo)
	arquivo, err := os.Open(nomeArquivo + ".csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
		return
	}

	arquivoJSON, err := os.Create(nomeArquivo + ".json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")
	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			if len(dados) > 1 {
				cidade := model.Cidade{}
				cidade.Nome = dados[0]
				cidade.Estado = dados[1]
				fmt.Printf("Cidade: %+v\r\n", cidade)
				cidadeJSON, err := json.Marshal(cidade)
				if err != nil {
					fmt.Println("[main] Houve um erro ao gerar o json do item ", item, ". Erro: ", err.Error())
				}
				escritor.WriteString("  " + string(cidadeJSON))
				if (indiceItem + 1) < len(linha) {
					escritor.WriteString(",\r\n")
				}
			} else {
				fmt.Println("Erro no nome da cidade: ", dados)
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
	fmt.Println(time.Now(), " - A tradução do arquivo: ", nomeArquivo, " foi finalizada")
	orquestrador.Done()
}
