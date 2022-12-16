package steps

import (
	"fmt"
	"testing"
)

var CountTableS3 = []struct {
	cnt int
}{
	{cnt: 10},
	{cnt: 1000},
	{cnt: 10000},
	{cnt: 100000},
	{cnt: 500000},
}

func BenchmarkConcurrentS3(b *testing.B) {
	for _, v := range CountTableS3 {
		b.Run(fmt.Sprintf("%d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessConcurrent(v.cnt)
			}
		})
	}
}
func BenchmarkIFanOut(b *testing.B) {
	for _, v := range CountTableS3 {
		b.Run(fmt.Sprintf("10 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				S3(10, v.cnt)
			}
		})
		b.Run(fmt.Sprintf("100 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				S3(100, v.cnt)
			}
		})
		b.Run(fmt.Sprintf("500 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				S3(500, v.cnt)
			}
		})
	}
}
func BenchmarkStatS3(b *testing.B) {
	for _, v := range CountTableS3 {
		b.Run(fmt.Sprintf("12 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessConcurrent(v.cnt)
				//S3(12, v.cnt)
			}
		})
		b.Run(fmt.Sprintf("144 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessConcurrent(v.cnt)
				//S3(144, v.cnt)
			}
		})
		b.Run(fmt.Sprintf("500 worker %d image", v.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessConcurrent(v.cnt)
				//S3(540, v.cnt)
			}
		})
	}
}
