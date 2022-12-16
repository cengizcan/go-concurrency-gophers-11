package image

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 50},
	{input: 500},
}

func BenchmarkReadTable(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ReadSeq(v.input)
			}
		})
	}
}

func BenchmarkRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadImage(fmt.Sprintf("IMG_%d", b.N))
	}
}
func BenchmarkReadWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := ReadImage(fmt.Sprintf("IMG_%d", b.N))
		WriteImage(img)
	}
}
