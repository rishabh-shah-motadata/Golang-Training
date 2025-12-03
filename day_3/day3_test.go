package day3

import "testing"

func twoAllocs() string {
	b := []byte{}
	b = append(b, 'a')
	c := []byte{}
	c = append(c, 'b')
	return string(b) + string(c)
}

func BenchmarkTwoAllocs(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = twoAllocs()
	}
}
