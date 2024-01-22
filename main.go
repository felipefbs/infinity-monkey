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

const monkeyNumber = 10000

func main() {
	// Reading text from file
	textByte, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.ReplaceAll(string(textByte), "\n", "")
	wordList := strings.Split(text, " ")

	// Creating pubsub topics and monkeys to find the words
	inc := pubsub.NewInc()
	returnChan := inc.Subscribe("return", 1)
	receiveChan := inc.Subscribe("words", 1)
	for i := 0; i < monkeyNumber; i++ {
		go Monkey(inc, i, receiveChan)
	}

	defer inc.Close()

	// Publishing all the words from text on the topic "words"
	found := make([]pubsub.Token, len(wordList))

	for index, word := range wordList {
		token := pubsub.Token{
			Index: index,
			Word:  word,
			Found: false,
		}
		found[index] = token
		inc.Publish("words", token)
	}

	// This anonm function is responsible to read all the found words
	go func() {
		for f := range returnChan {
			found[f.Index] = f
		}
	}()

	// The main loop where checks if all words were found or print how many are found already
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
