package fibonacci

import (
	"strconv"
	"testing"
)

// sink para evitar que o compilador otimize as chamadas
var sinkInt int

// -------- Benchmarks para Fibonacci (int) --------

func BenchmarkFibonacci_Sub(b *testing.B) {
	cases := []int{10, 30, 50, 92} // 92 é o máx. seguro em int64
	for _, n := range cases {
		b.Run("N="+strconv.Itoa(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sinkInt = Fibonacci(n)
			}
		})
	}
}

// Opcional: benchmark de geração da sequência (usa Sequence "int", não a versão big)
func BenchmarkSequence_Sub(b *testing.B) {
	cases := []int{10, 30, 50}
	for _, n := range cases {
		b.Run("N="+strconv.Itoa(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = Sequence(n)
			}
		})
	}
}
