package steps

import (
	"sync"

	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

type Stats struct {
	successCnt, failCnt, readBytes, writtenBytes int
}

func FanOut(input <-chan string, workerCnt int, fn ProcessImage) <-chan *ProcessResponse {
	num := numerator.NewV2()

	output := make(chan *ProcessResponse)
	var wg sync.WaitGroup

	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			for {
				path, ok := <-input

				if !ok {
					return // return when closed
				}
				output <- fn(path, num)
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func Process(workerCnt, imageCnt int, fn ProcessImage) *Stats {
	input := make(chan string)

	output := FanOut(input, workerCnt, fn)
	go func() {
		for p := range image.ScanImages(imageCnt) {
			input <- p
		}
		close(input)
	}()

	return StatGen(output)

}

func StatGen(output <-chan *ProcessResponse) *Stats {
	var successCnt, failCnt, readBytes, writtenBytes int
	for r := range output {
		if r.err != nil {
			failCnt++
		} else {
			successCnt++
		}
		readBytes += r.Read
		writtenBytes += r.Written
	}

	return &Stats{
		successCnt:   successCnt,
		failCnt:      failCnt,
		readBytes:    readBytes,
		writtenBytes: writtenBytes,
	}
}
func S3(workerCnt, imageCnt int) *Stats {
	return Process(workerCnt, imageCnt, ProcessImageS2)
}
