package steps

import (
	"context"
	"sync"

	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

func fanIn(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup // <2>
	multiplexedStream := make(chan string)

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan string) {
			defer wg.Done()
			for i := range c {
				multiplexedStream <- i
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func S7(ctx context.Context, workerCnt int, imageCnt []int) *Stats {
	num := numerator.NewV2()

	wp := New[*ProcessResponse](workerCnt)
	go wp.Run(ctx)

	go func() {
		channels := make([]<-chan string, len(imageCnt))
		for i, cnt := range imageCnt {
			channels[i] = image.ScanImages(cnt)
		}

		for p := range fanIn(channels...) {
			wp.Tasks <- NewTask(p, ProcessImageS5, num)
		}
		close(wp.Tasks)
	}()

	return StatGen(wp.Results)

}
