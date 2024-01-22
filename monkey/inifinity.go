package monkey

import (
	"crypto/rand"
	"math/big"
)

func CharacterPicker() string {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(5011)))
	x = x.Add(x, big.NewInt(int64(32)))

	return string(rune(uint32(x.Uint64())))
}

func TrueInfinityMonkey(word string) string {
	i := 0
	found := ""
	for {
		letter := CharacterPicker()
		if i == len(word) {
			break
		}
		currentLetter := word[i]
		if string(currentLetter) == letter {
			found += string(currentLetter)
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
		letter := CharacterPicker()
		if i == len(word) {
			break
		}
		currentLetter := word[i]
		if string(currentLetter) == letter {
			found += string(currentLetter)
			i++
		}

		if len(found) == len(word) {
			return found
		}
	}

	return ""
}
