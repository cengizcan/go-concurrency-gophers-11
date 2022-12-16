package steps

import (
	"fmt"
	"testing"
)

var CountTable = []struct {
	cnt int
}{
	{cnt: 10},
	{cnt: 100},
	{cnt: 1000},
	{cnt: 10000},
}

func BenchmarkSequential(b *testing.B) {
	for _, v := range CountTable {
		b.Run(fmt.Sprintf("%d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessSequential(v.cnt)
			}
		})
	}
}
func BenchmarkConcurrent(b *testing.B) {
	for _, v := range CountTable {
		b.Run(fmt.Sprintf("%d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessConcurrent(v.cnt)
			}
		})
	}
}

func BenchmarkStatS2(b *testing.B) {
	for _, v := range CountTable {
		b.Run(fmt.Sprintf("%d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//ProcessSequential(v.cnt)
				ProcessConcurrent(v.cnt)
			}
		})
	}
}
