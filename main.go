package main

import (
	"fmt"

	. "markov.chains/markov"
	. "markov.chains/tokenizer"
)

func main() {
	d := `abcdefga`

	t := NewTokenizer(d)
	m := NewModel(2, t)

	m.Train(d)

	prompt := "a"

	for i := 0; i < 10; i++ {
		sm, err := m.Forward(prompt)
		if err != nil {
			panic(err)
		}

		st, err := sm.Sample(0.8)
		if err != nil {
			panic(err)
		}

		d := t.Decode(st)

		fmt.Printf("Sample from\n")
		sm.Print(t)
		fmt.Printf("Chose: %s\n", d)

		prompt += d
	}
}
