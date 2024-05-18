package markov

import (
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

func (m *Model) Predict(prompt string) (string, error) {
	enc := m.tokenizer.Encode(prompt)
	tok := enc[len(enc)-1]

	sm, err := m.layers[0].Row(int(tok))
	if err != nil {
		return "", err
	}

	if len(m.layers) > 1 {
		for i := 1; i < len(m.layers); i++ {
			layer := m.layers[i]
			nr, err := layer.Row(int(tok))
			if err != nil {
				return "", err
			}

			err = sm.Nudge(nr, i)
			if err != nil {
				return "", err
			}
		}
	}

	p, err := sm.Sample(0)
	if err != nil {
		return "", err
	}

	return m.tokenizer.Decode(p), nil
}
