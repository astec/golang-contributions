package gocontrib

import (
	"strconv"
	"testing"
)

func TestAtoiImproved(t *testing.T) {
	v, err := AtoiImproved("+12")
	t.Logf("v=%v, err=%v", v, err)

	v, err = AtoiImproved("-12")
	t.Logf("v=%v, err=%v", v, err)
}

func BenchmarkAtoi(b *testing.B) {

	benchmark := func(b *testing.B, v string, f func(string) (int, error)) {
		for i := 0; i < b.N; i++ {
			i, _ := f(v)
			BenchSink += i
		}
	}

	compare := func(name, v string) {
		for i := 1; i < 2; i++ {
			b.Run(name+"_original_"+strconv.Itoa(i), func(b *testing.B) {
				benchmark(b, v, strconv.Atoi)
			})
			b.Run(name+"_improved_"+strconv.Itoa(i), func(b *testing.B) {
				benchmark(b, v, AtoiImproved)
			})
		}
	}

	compare("unsigned", "12")
	compare("positive", "+12")
	compare("negative", "-12")
	compare("minus_only", "-")
	compare("plus_only", "+")

	b.Log(BenchSink)
}

var BenchSink int
var lastErr error
