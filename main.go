package main

import (
	"fmt"

	. "markov.chains/markov"
	. "markov.chains/tokenizer"
)

func main() {
	d := "The first one to say the thing is a fool"
	//	b, err := os.ReadFile("./dataset/inp.txt")
	//	if err != nil {
	//		panic(err)
	//	}

	// d := string(b)

	t := NewTokenizer(d)
	m := NewModel(2, t)

	m.Train(d)
	m.Print()

	prompt := "The"

	fmt.Print(prompt)

	for i := 0; i < 30; i++ {
		s, err := m.Predict(prompt)
		if err != nil {
			panic(err)
		}

		prompt += s

		fmt.Print(s)
	}
}
