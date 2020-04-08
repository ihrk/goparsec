package goparsec

import "testing"

var td = []byte(`99999866`)

func TestManyC(t *testing.T) {
	p := Many(Rune('9'))
	n, err := p.Parse(td)
	if err != nil {
		t.Error(err)
	}
	if n != 5 {
		t.Error("failed to parse")
	}
}

func BenchmarkRune(b *testing.B) {
	p := Many(Rune('9'))
	for i := 0; i < b.N; i++ {
		p.Parse(td)
	}
}
