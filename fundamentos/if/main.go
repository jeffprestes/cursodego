package main

import "fmt"

func main() {

	meses := 0
	situacao := true
	cidade := "Teste"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste")
	} else if cidade == "Athenas" {
		fmt.Println("O cliente vive em Athenas.")
	} else {
		fmt.Println("O cliente não vive no estado Teste")
	}

	/*
		descricao, status := haQuantoTempoEstaDevendo(meses)
		if status {

		}
	*/
	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	//A instrucao abaixo não compila
	//fmt.Println("Me passa o status?", descricao)

	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo."
		return
	}
	descricao = "O cliente esta com o carnê em dia."
	return
}
