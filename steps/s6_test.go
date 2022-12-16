package steps

import (
	"context"
	"fmt"
	"testing"
)

var CountTableS6 = []struct {
	cnt int
}{
	{cnt: 10000},
	{cnt: 500000},
	{cnt: 750000},
}

func BenchmarkStatS6(b *testing.B) {
	workerCount := []int{100, 500, 1000}
	ctx := context.TODO()
	for _, img := range CountTableS6 {
		for _, c := range workerCount {
			b.Run(fmt.Sprintf("%d worker %d image", c, img.cnt), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					//S6Concurrent(img.cnt)
					S6(ctx, c, img.cnt)
				}
			})
		}
	}
}

func BenchmarkStatConcurrent(b *testing.B) {
	//ctx := context.TODO()
	for _, img := range CountTableS6 {
		b.Run(fmt.Sprintf("%d image", img.cnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				S6Concurrent(img.cnt)
			}
		})
	}
}
