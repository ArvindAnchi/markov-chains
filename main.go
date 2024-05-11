package main

import (
	"fmt"
	"log"
	"os"

	. "markov.chains/tokenizer"
)

func main() {
	dat, err := os.ReadFile("./dataset/inp.txt")
	if err != nil {
		log.Fatal(err)
	}

	t := NewTokenizer(string(dat))

	t.Print()

	fmt.Println(t.Encode("Kria's"))
	fmt.Println(t.Decode(923))
}
