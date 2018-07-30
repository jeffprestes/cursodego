package matematica

import (
	"testing"
)

func TestMultiplicador(t *testing.T) {
	total := Multiplicador(2, 3)
	if total != 6 {
		t.Errorf("Erro na funcao multiplicador")
	}
}
