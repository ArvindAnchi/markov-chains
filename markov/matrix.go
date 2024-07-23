package markov

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"

	. "markov.chains/tokenizer"
)

type Matrix struct {
	Rows    int
	Cols    int
	startId uint16
	es      []uint16
}

func NewMat(rows, cols int) *Matrix {
	es := make([]uint16, rows*cols)

	return &Matrix{
		Rows: rows,
		Cols: cols,
		es:   es,
	}
}

func (m *Matrix) idxOf(row, col int) (int, error) {
	if col > m.Cols {
		return 0, errors.New(fmt.Sprintf("MAT_IDOF: Column %d out of bounds for matrix with %d columns", col, m.Cols))
	}
	if row > m.Rows {
		return 0, errors.New(fmt.Sprintf("MAT_IDOF: Row %d out of bounds for matrix with %d rows", row, m.Rows))
	}

	return row*m.Cols + col, nil
}

func (m *Matrix) Print(t *Tokenizer) {
	p := make([][2]int, m.Rows)

	for i := 0; i < m.Rows; i++ {
		p[i] = [2]int{i, int(m.es[i])}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i][0] > p[i][1]
	})

	for i := 0; i < len(p); i++ {
		fmt.Printf("%s: %d", t.Decode(uint16(p[i][0])), p[i][1])
	}
}

func (m *Matrix) Fill(value uint16) {
	for i := 0; i < len(m.es); i++ {
		m.es[i] = value
	}
}

func (m *Matrix) Nudge(m1 *Matrix, r int) error {
	if m.Cols != m1.Cols {
		return errors.New(fmt.Sprintf("MAT_SUM: Incompatable shapes m:(%d %d) and m1:(%d %d)", m.Rows, m.Cols, m1.Rows, m1.Cols))
	}
	if m.Rows != m1.Rows {
		return errors.New(fmt.Sprintf("MAT_SUM: Incompatable shapes m:(%d %d) and m1:(%d %d)", m.Rows, m.Cols, m1.Rows, m1.Cols))
	}
	if len(m.es) != len(m1.es) {
		return errors.New(fmt.Sprintf("MAT_SUM: Incompatable inner array sizes m:(%d) and m1:(%d)", len(m.es), len(m1.es)))
	}

	for i := 0; i < len(m.es); i++ {
		m.es[i] += m1.es[i] / uint16(r)
	}

	return nil
}

func (m *Matrix) Inc(row, col int) error {
	i, err := m.idxOf(row, col)
	if err != nil {
		return err
	}

	if m.es[i] < 65535 {
		m.es[i] += 1
	} else {
		fmt.Printf("overflow %d %d\n", row, col)
	}

	return nil
}

func (m *Matrix) Sample(topp float64) (uint16, error) {
	if m.Rows != 1 {
		return 0, errors.New(fmt.Sprintf("MAT_SAMPLE: Expected matrix with 1 row got %d rows", m.Rows))
	}

	t := 0
	sMax := float64(slices.Max(m.es))

	for _, tp := range m.es {
		p := float64(tp) / sMax
		if p < topp {
			continue
		}

		t += int(tp)
	}

	if t == 0 {
		return 0, errors.New("MAT_SAMPLE: Unable to sample from 0 sum probs")
	}

	r := rand.IntN(t)

	for i, tp := range m.es {
		p := float64(tp) / sMax
		if p < topp {
			continue
		}

		r -= int(tp)

		if r < 0 {
			return uint16(i), nil
		}
	}

	return 0, errors.New(fmt.Sprintf("MAT_SAMPLE: Reached unreachable (r=%d)", r))
}

func (m *Matrix) Row(rowIdx int) (*Matrix, error) {
	sIdx, err := m.idxOf(rowIdx, 0)
	eIdx := sIdx + m.Cols

	if err != nil {
		return nil, err
	}

	return &Matrix{
		Rows:    1,
		Cols:    m.Cols,
		startId: uint16(rowIdx),
		es:      m.es[sIdx:eIdx],
	}, nil
}
