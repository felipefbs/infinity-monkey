package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func main() {
	for {
		dist := make(map[string]int)
		var l []string
		for i := 0; i < 26; i++ {
			x, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
			a := alphabet[x.Uint64()]
			dist[a]++
			l = append(l, a)
		}

		sum := 0
		for _, v := range dist {
			sum += v
		}

		fmt.Println(l)
		fmt.Println(dist)
		x := float64(sum) / float64(len(dist))
		if x == 1 {
			break
		}
		fmt.Println()
	}
}
