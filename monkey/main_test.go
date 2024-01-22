package monkey_test

import (
	"fmt"
	"testing"

	"github.com/felipefbs/infinity-monkey/monkey"
)

func BenchmarkLetterPicker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monkey.CharacterPicker()
	}
}

var table = []string{ /*"a", "aa" "aaa", "aaaa", "aaaaa"*/ }

func BenchmarkEasy(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprint(v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				monkey.EasyInfinityMonkey(v)
			}
		})
	}
}

func BenchmarkTrue(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprint(v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				monkey.TrueInfinityMonkey(v)
			}
		})
	}
}
