package markov

import (
	"errors"
	"fmt"
	"math/rand/v2"

	. "markov.chains/tokenizer"
)

type Matrix struct {
	Rows int
	Cols int
	es   []uint16
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
	fmt.Print(m.es)
	fmt.Print("[\n\t")

	for j := 0; j < m.Cols; j++ {
		fmt.Printf("%s:\t", t.Decode(uint16(j)))
	}

	fmt.Print("\n")

	for i := 0; i < m.Rows; i++ {
		fmt.Printf("%s:", t.Decode(uint16(i)))

		for j := 0; j < m.Cols; j++ {
			idx, err := m.idxOf(i, j)

			if err != nil {
				panic(err)
			}

			fmt.Printf("\t%d", m.es[idx])
		}

		fmt.Printf("\n")
	}

	fmt.Print("]\n")
}

func (m *Matrix) Fill(value uint16) {
	for i := 0; i < len(m.es); i++ {
		m.es[i] = value
	}
}

func (m *Matrix) Sum(m1 *Matrix) error {
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
		m.es[i] += m1.es[i]
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

func (m *Matrix) Sample(rowIdx int) (uint16, error) {
	sIdx, err := m.idxOf(rowIdx, 0)
	eIdx := sIdx + m.Cols

	if err != nil {
		return 0, err
	}

	t := 0

	for _, tp := range m.es[sIdx:eIdx] {
		t += int(tp)
	}

	if t == 0 {
		return uint16(rand.IntN(m.Cols)), nil
	}

	r := rand.IntN(t)

	for i, tp := range m.es[sIdx:eIdx] {
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
		Rows: 1,
		Cols: m.Cols,
		es:   m.es[sIdx:eIdx],
	}, nil
}
