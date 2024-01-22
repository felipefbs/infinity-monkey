package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/felipefbs/infinity-monkey/monkey"
	"github.com/felipefbs/infinity-monkey/pubsub"
)

func main() {
	inc := pubsub.NewInc()
	monkeyNumber := 10000

	textByte, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := strings.ReplaceAll(string(textByte), "\n", "")

	wordList := strings.Split(text, " ")

	returnChan := inc.Subscribe("return", len(wordList))
	receiveChan := inc.Subscribe("book", len(wordList))
	for i := 0; i < monkeyNumber; i++ {
		go Monkey(inc, i, receiveChan)
	}

	found := make([]pubsub.Token, len(wordList))

	for index, word := range wordList {
		token := pubsub.Token{
			Index: index,
			Word:  word,
			Found: false,
		}
		found[index] = token
		inc.Publish("book", token)
	}

	go func() {
		for f := range returnChan {
			found[f.Index] = f
		}
	}()

	countFoundChan := make(chan int)
	go CountFound(&found, countFoundChan)

	aux := 0
	for {
		aux = <-countFoundChan
		if aux != <-countFoundChan {
			fmt.Printf("found %v out of %v\n", <-countFoundChan, len(wordList))
		}
		if <-countFoundChan == len(wordList) {
			break
		}
	}

	fmt.Println(found)
}

func CountFound(found *[]pubsub.Token, count chan int) {
	for {
		counter := 0
		for _, t := range *found {
			if t.Found {
				counter++
			}
		}

		count <- counter
	}
}

func PrintNotFound(found *[]pubsub.Token) {
	for {
		time.Sleep(10 * time.Second)
		for _, t := range *found {
			if !t.Found {
				fmt.Println(t.Word)
			}
		}
	}
}

func Monkey(pub *pubsub.Publisher, id int, receiveChan <-chan pubsub.Token) {
	for {
		toFind := <-receiveChan
		found := monkey.TrueInfinityMonkey(toFind.Word)
		fmt.Println("monkey", id, "found word:", found)
		toFind.Found = true
		pub.Publish("return", toFind)
	}
}
