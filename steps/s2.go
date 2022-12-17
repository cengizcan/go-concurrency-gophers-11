package steps

import (
	"sync"

	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

type ProcessResponse struct {
	Name    string
	Meta    map[string]string
	Id      int
	Read    int
	Written int
	err     error
}

type ProcessImage func(name string, num numerator.Sequential) *ProcessResponse

func ProcessImageS2(name string, num numerator.Sequential) *ProcessResponse {
	// read
	img := image.ReadImage(name)
	r := len(img.Bytes)
	// resize
	img = image.Resize(img)
	// generate doc number
	img.Id = num.Next()
	// save
	image.WriteImage(img)

	return &ProcessResponse{
		Name:    img.Name,
		Meta:    img.Meta,
		Id:      img.Id,
		Read:    r,
		Written: len(img.Bytes),
	}
}

// Sequential implementation
func ProcessSequential(cnt int) {
	for path := range image.ScanImages(cnt) {
		ProcessImageS2(path, numerator.NewV1())
	}
}

// Concurrent implementation
func ProcessConcurrent(cnt int) {
	wg := new(sync.WaitGroup)
	for p := range image.ScanImages(cnt) {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			ProcessImageS2(path, numerator.NewV2())
		}(p)
	}
	wg.Wait()
}
