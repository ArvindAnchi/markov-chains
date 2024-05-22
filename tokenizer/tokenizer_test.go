package tokenizer

import "testing"

func TestTokens(t *testing.T) {
	p := "aabbca"
	tk := NewTokenizer(p)
	v := *tk.Vocab()

	if _, ok := v["a"]; !ok {
		t.Fatalf("'a' Not in vocab")
	}
	if _, ok := v["b"]; !ok {
		t.Fatalf("'a' Not in vocab")
	}
	if _, ok := v["c"]; !ok {
		t.Fatalf("'a' Not in vocab")
	}
}

func TestEncode(t *testing.T) {
	p := "aabbca"
	tk := NewTokenizer(p)

	if tks := tk.Encode("a"); tks[0] != 0 {
		t.Fatalf("Encode 'a' expected 0 got %d", tks[0])
	}
}

func TestDecode(t *testing.T) {
	p := "aabbca"
	tk := NewTokenizer(p)

	if tks := tk.Decode(0); tks != "a" {
		t.Fatalf("Decode 0 expected 'a' got '%s'", tks)
	}
}
