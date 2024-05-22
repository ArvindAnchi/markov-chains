package markov

import "testing"

func TestMatFill(t *testing.T) {
	m := NewMat(2, 2)

	m.Fill(2)

	for i, v := range m.es {
		if v != 2 {
			t.Fatalf("MAT_FILL: Expected 2 got %d at %d", v, i)
		}
	}
}

func TestIdxOf(t *testing.T) {
	m := NewMat(2, 3)
	m.es = []uint16{
		1, 2, 3,
		4, 5, 6,
	}

	t.Log(m.es)

	val1, err := m.idxOf(0, 0)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val1] != 1 {
		t.Fatalf("MAT_IDX_OF: Expected 1 got %d at (%d, %d)", m.es[val1], 0, 0)
	}

	val2, err := m.idxOf(0, 1)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val2] != 2 {
		t.Fatalf("MAT_IDX_OF: Expected 2 got %d at (%d, %d)", m.es[val2], 0, 1)
	}

	val3, err := m.idxOf(0, 2)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val3] != 3 {
		t.Fatalf("MAT_IDX_OF: Expected 3 got %d at (%d, %d)", m.es[val3], 1, 0)
	}

	val4, err := m.idxOf(1, 0)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val4] != 4 {
		t.Fatalf("MAT_IDX_OF: Expected 4 got %d at (%d, %d)", m.es[val4], 1, 0)
	}

	val5, err := m.idxOf(1, 1)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val5] != 5 {
		t.Fatalf("MAT_IDX_OF: Expected 5 got %d at (%d, %d)", m.es[val5], 1, 1)
	}

	val6, err := m.idxOf(1, 2)

	if err != nil {
		t.Fatalf("MAT_IDX_OF: %v", err)
	}
	if m.es[val6] != 6 {
		t.Fatalf("MAT_IDX_OF: Expected 6 got %d at (%d, %d)", m.es[val6], 1, 2)
	}
}

func TestInc(t *testing.T) {
	m := NewMat(2, 3)
	m.es = []uint16{
		1, 2, 3,
		4, 5, 6,
	}

	err := m.Inc(1, 2)

	if err != nil {
		t.Fatalf("MAT_INC: %v", err)
	}

	idx, err := m.idxOf(1, 2)

	if err != nil {
		t.Fatalf("MAT_INC: %v", err)
	}
	if m.es[idx] != 7 {
		t.Fatalf("MAT_INC: Expected 7 at (1, 2) got %d", m.es[idx])
	}
}

func TestRow(t *testing.T) {
	m := NewMat(3, 3)
	m.es = []uint16{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	r1, err := m.Row(0)
	if err != nil {
		t.Fatalf("MAT_ROW: %v", err)
	}

	for idx, v := range [3]uint16{1, 2, 3} {
		if r1.es[idx] != v {
			t.Fatalf("MAT_ROW: Expected %d got %d", v, r1.es[idx])
		}
	}

	r2, err := m.Row(1)
	if err != nil {
		t.Fatalf("MAT_ROW: %v", err)
	}

	for idx, v := range [3]uint16{4, 5, 6} {
		if r2.es[idx] != v {
			t.Fatalf("MAT_ROW: Expected %d got %d", v, r2.es[idx])
		}
	}

	r3, err := m.Row(2)
	if err != nil {
		t.Fatalf("MAT_ROW: %v", err)
	}

	for idx, v := range [3]uint16{7, 8, 9} {
		if r3.es[idx] != v {
			t.Fatalf("MAT_ROW: Expected %d got %d", v, r3.es[idx])
		}
	}
}
