package main

import (
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

	prompt := "Th"

	sm, err := m.Forward(prompt)
	if err != nil {
		panic(err)
	}

	sm.Print(t)
}
