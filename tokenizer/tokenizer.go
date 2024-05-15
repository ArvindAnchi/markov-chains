package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type Tokenizer struct {
	vocab map[string]uint16
}

var re = regexp.MustCompile(`(?: ?[^\s\w\d]|(?: ?\w+|\d+))`)

func NewTokenizer(corpus string) *Tokenizer {
	corpus = strings.ToLower(corpus)

	toks := re.FindAll([]byte(corpus), -1)

	v := make(map[string]uint16)
	i := 0

	for _, tok := range toks {
		if _, ok := v[string(tok)]; !ok {
			v[string(tok)] = uint16(i)
			i++
		}
	}

	return &Tokenizer{
		vocab: v,
	}
}

func (t *Tokenizer) Vocab() *map[string]uint16 {
	return &t.vocab
}

func (t *Tokenizer) Print() {
	for tok, i := range t.vocab {
		fmt.Printf("'%s' -> %d\n", tok, i)
	}
}

func (t *Tokenizer) Encode(text string) []uint16 {
	text = strings.ToLower(text)

	splits := re.FindAll([]byte(text), -1)
	tokens := make([]uint16, len(splits))

	for i, tok := range splits {
		if tokId, ok := t.vocab[string(tok)]; ok {
			tokens[i] = tokId
		} else {
			tokens[i] = 65535
		}
	}

	return tokens
}

func (t *Tokenizer) Decode(tok uint16) string {
	for vtok, i := range t.vocab {
		if i == tok {
			return vtok
		}
	}

	return "<UNK>"
}
