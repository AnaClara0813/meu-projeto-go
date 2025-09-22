// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

import (
	"fmt"//fmt é um pacote da linguagem Go usado para formatar e imprimir dados.
//Ele é equivalente ao print em outras linguagens, mas muito mais poderoso, porque permite formatar strings, números, variáveis etc.
	
	// "github.com/AnaClara0813/meu-projeto-go/internal/fibonacci"
	// "github.com/AnaClara0813/meu-projeto-go/internal/hello"
	//"github.com/AnaClara0813/meu-projeto-go/internal/memoria-go"
	//"github.com/AnaClara0813/meu-projeto-go/internal/saudacoes"
	"github.com/AnaClara0813/meu-projeto-go/internal/anamnese"
)

func main() {
	// fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
	// hello.SayHello()
	// n := 8
	// valor := fibonacci.Fibonacci(n)
	// fmt.Printf("F(%d) = %d\n", n, valor)
	// fibonacci.PrintSequence(n)

	//memoriago.Play()

	//saudar := saudacoes.NovaSaudacao("Bem-vinda")
	//fmt.Println(saudar("Ana Clara")) // -> "Bem-vinda, Ana Clara! "

	//saudacoes.PrintSaudacao("FacInpro")

	nome := "Ana"  // string
	idade := 27    // int
	peso := 62.0   // float64
	altura := 1.62 // float64

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)

	imc := anamnese.Resultado(peso, altura)
	fmt.Printf("Resultado (IMC): %.2f\n", imc)

}
