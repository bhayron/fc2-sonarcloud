package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: resultado %d. Esperado: %d", total, 30)
	}
}

//ci com golang e github
