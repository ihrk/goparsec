package goparsec

import (
	"regexp"
	"testing"
)

var data = []byte(` 



<sd>
<sa> 
tgfgg

<gfd>
sdfdfgdf
<fd>
sdad
asdasdf
sdfsdf
sdf

`)

var tagFinder = And(
	Rune('<'),
	Some(RuneFunc(func(r rune) bool { return r != '>' })),
	Rune('>'),
)

var reg = regexp.MustCompile(`<[^>]*>`)

func TestFind(t *testing.T) {
	res1 := FindAllString(data, tagFinder)
	if len(res1) != 4 {
		t.Error(res1)
	}
	res2 := reg.FindAllString(string(data), -1)
	if len(res2) != 4 {
		t.Error(res2)
	}
	if res1[1] != res2[1] {
		t.Error(res1, res2)
	}
}

func BenchmarkFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindAll(data, tagFinder)
	}
}

func BenchmarkRegexp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg.FindAll(data, -1)
	}
}
