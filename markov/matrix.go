package markov

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Matrix struct {
	Rows int
	Cols int
	es   []uint8
}

func NewMat(rows, cols int) *Matrix {
	es := make([]uint8, rows*cols)

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
		return 0, errors.New(fmt.Sprintf("MAT_IDOF: Row %d out of bounds for matrix with %d rows", col, m.Cols))
	}

	return row * col, nil
}

func (m *Matrix) Inc(row, col int) error {
	i, err := m.idxOf(row, col)

	if err != nil {
		return err
	}

	m.es[i] += 1

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

	r := rand.IntN(t)

	for i, tp := range m.es[sIdx:eIdx] {
		if r < 0 {
			return uint16(i), nil
		}

		r -= int(tp)
	}

	return 0, errors.New("MAT_SAMPLE: Reached unreachable")
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
