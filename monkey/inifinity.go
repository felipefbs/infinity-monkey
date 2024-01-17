package monkey

import (
	"crypto/rand"
	"math/big"
)

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func LetterPicker() string {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
	return alphabet[x.Uint64()]
}

func TrueInfinityMonkey(word string) string {
	i := 0
	found := ""
	for {
		// time.Sleep(1 / 2 * time.Second)
		letter := LetterPicker()
		if i == len(word) {
			break
		}
		letraDaVez := word[i]
		if string(letraDaVez) == letter {
			found += string(letraDaVez)
			i++
		} else {
			found = ""
			i = 0
		}
		if len(found) == len(word) {
			return found
		}
	}

	return ""
}

func EasyInfinityMonkey(word string) string {
	i := 0
	found := ""
	for {
		// time.Sleep(1 / 2 * time.Second)
		letter := LetterPicker()
		if i == len(word) {
			break
		}
		letraDaVez := word[i]
		if string(letraDaVez) == letter {
			found += string(letraDaVez)
			i++
		}

		if len(found) == len(word) {
			return found
		}
	}

	return ""
}
