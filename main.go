// Nome do pacote
package main

// Importa o fmt
import "fmt"

// função main realiza a soma jogando o valor na variavel
func main() {

	vsoma := soma(5,5)

	// printa variável na tela
	fmt.Println(vsoma)

}

// efetua a soma de dois inteiros retornando inteiro
func soma (a , b int ) int {
	return a + b
}