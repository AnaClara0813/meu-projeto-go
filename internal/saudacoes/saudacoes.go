package saudacoes

import "fmt"


func NovaSaudacao(prefixo string) func(string) string {
	return func(nome string) string {
		return fmt.Sprintf("%s, %s! ", prefixo, nome)
	}
}


func PrintSaudacao(nome string) {
	func(n string) {
		fmt.Printf("Olá, %s! Bem-vinda a função anonima Go! 😎\n", n)
	}(nome)
}
