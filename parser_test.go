package goparsec

import (
	"testing"
	"unicode"
)

var td = []byte("false")

var ttd = []byte("2")

func BenchmarkRune(b *testing.B) {
	p := Rune('f')
	for i := 0; i < b.N; i++ {
		_, _ = p(td)
	}
}

func TestStr(t *testing.T) {
	p := Str("false")
	_, err := p(td)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkStr(b *testing.B) {
	p := Str("false")
	for i := 0; i < b.N; i++ {
		_, _ = p(td)
	}
}

func TestTrimFunc(t *testing.T) {
	p1 := Many(Func(unicode.IsDigit))
	p2 := TrimFunc(unicode.IsDigit)
	n1, err1 := p1(ttd)
	n2, err2 := p2(ttd)
	if n1 != n2 || n2 == 0 {
		t.Error(err1, err2)
	}
}

func BenchmarkManyFunc(b *testing.B) {
	p := Many(Func(unicode.IsDigit))
	for i := 0; i < b.N; i++ {
		_, _ = p(ttd)
	}
}

func BenchmarkTrimFunc(b *testing.B) {
	p := TrimFunc(unicode.IsDigit)
	for i := 0; i < b.N; i++ {
		_, _ = p(ttd)
	}
}
