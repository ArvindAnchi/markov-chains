package main

import (
	"fmt"
	"log"

	. "markov.chains/markov"
	. "markov.chains/tokenizer"
)

func main() {
	d := "The first one to say the thing is a fool"
	t := NewTokenizer(d)

	vc := len(*t.Vocab())
	toks := t.Encode(d)

	fmt.Printf("Vocab length = %d\n", vc)

	mn1 := NewMat(vc, vc)
	mn2 := NewMat(vc, vc)

	for i := 2; i < len(toks); i += 2 {
		err := mn2.Inc(int(toks[i-2]), int(toks[i]))

		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 1; i < len(toks); i++ {
		err := mn1.Inc(int(toks[i-1]), int(toks[i]))

		if err != nil {
			log.Fatal(err)
		}
	}

	mn1.Print()
	mn2.Print()

	prompt := "The"

	fmt.Print(prompt)

	for i := 0; i < 10; i++ {
		tok := t.Encode(prompt)[0]

		sm := NewMat(1, vc)
		sm.Fill(0)

		n1r, err := mn1.Row(int(tok))
		if err != nil {
			log.Fatal(err)
		}

		n2r, err := mn2.Row(int(tok))
		if err != nil {
			log.Fatal(err)
		}

		err = sm.Sum(n1r)
		if err != nil {
			log.Fatal(err)
		}

		err = sm.Sum(n2r)
		if err != nil {
			log.Fatal(err)
		}

		p, err := mn1.Sample(int(tok))
		if err != nil {
			log.Fatal(err)
		}

		s := t.Decode(p)

		prompt += s

		fmt.Print(s)
	}
}
