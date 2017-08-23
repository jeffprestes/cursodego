package main

import "fmt"
import "github.com/jeffprestes/cursodego/funcoes_basico/matematica"

func main() {

	resultado := matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicacao foi: %d\r\n", resultado)
	
}
