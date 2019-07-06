// Nome do pacote
package main

// Importa o pacote testing
import "testing"

// Funçaõ de teste
func TestSoma(t *testing.T) {

	resultado := soma(5,5)

	if resultado != 10 {
		t.Errorf("A soma está incorreta, retornou %d, quando deveria retornar %d.", resultado, 10)
	}

	
}