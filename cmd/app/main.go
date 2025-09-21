// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"github.com/AnaClara0813/meu-projeto-go/internal/fibonacci"
	"github.com/AnaClara0813/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	hello.SayHello()

	n := 8
	
	valor := fibonacci.Fibonacci(n)

	fmt.Printf("F(%d) = %d\n", n, valor)

	fibonacci.PrintSequence(n)
}
