// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

import (
	"fmt"
	// "github.com/AnaClara0813/meu-projeto-go/internal/fibonacci"
	// "github.com/AnaClara0813/meu-projeto-go/internal/hello"
	//"github.com/AnaClara0813/meu-projeto-go/internal/memoria-go"
	"github.com/AnaClara0813/meu-projeto-go/internal/saudacoes"

)

func main() {
	// fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
	// hello.SayHello()
	// n := 8
	// valor := fibonacci.Fibonacci(n)
	// fmt.Printf("F(%d) = %d\n", n, valor)
	// fibonacci.PrintSequence(n)

	//memoriago.Play()

	saudar := saudacoes.NovaSaudacao("Bem-vinda")
	fmt.Println(saudar("Ana Clara")) // -> "Bem-vinda, Ana Clara! "

	saudacoes.PrintSaudacao("FacInpro")



}
