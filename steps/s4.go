package steps

import (
	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

func ProcessImageS4(name string, num numerator.Sequential) *ProcessResponse {
	// read
	img := image.ReadImage(name)
	r := len(img.Bytes)
	// resize
	img = image.Resize(img)
	// generate doc number
	img.Id = num.Next()
	// compress
	comp := image.Compress(img.Bytes)
	img.Bytes = comp.Output
	img.Meta["Compressor"] = comp.SvcName
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
func S4(workerCnt, imageCnt int) *Stats {
	return Process(workerCnt, imageCnt, ProcessImageS4)
}
