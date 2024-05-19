package markov

import (
	"errors"

	. "markov.chains/tokenizer"
)

type Model struct {
	tokenizer *Tokenizer

	layers []Matrix
}

func NewModel(layerCount int, tokenizer *Tokenizer) *Model {
	l := make([]Matrix, layerCount)
	v := len(*tokenizer.Vocab())

	for i := 0; i < layerCount; i++ {
		l[i] = *NewMat(v, v)
	}

	return &Model{
		tokenizer: tokenizer,
		layers:    l,
	}
}

func (m *Model) Print() {
	for _, mat := range m.layers {
		mat.Print(m.tokenizer)
	}
}

func (m *Model) Train(data string) error {
	toks := m.tokenizer.Encode(data)

	for l, layer := range m.layers {
		n := l + 1

		for i := n; i < len(toks); i++ {
			err := layer.Inc(int(toks[i-n]), int(toks[i]))

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Model) Forward(prompt string) (*Matrix, error) {
	toks := m.tokenizer.Encode(prompt)
	lenc := len(toks) - 1

	if lenc < 0 {
		return nil, errors.New("MODEL_FORWARD: Got 0 tokens")
	}

	sm, err := m.layers[0].Row(int(toks[lenc]))
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(m.layers); i++ {
		if i > lenc {
			break
		}

		nr, err := m.layers[i].Row(int(toks[lenc-i]))
		if err != nil {
			return nil, err
		}

		err = sm.Nudge(nr, i+1)
		if err != nil {
			return nil, err
		}
	}

	return sm, nil
}
