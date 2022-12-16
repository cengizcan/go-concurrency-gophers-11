package steps

import (
	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

func ProcessImageS5(name string, num numerator.Sequential) *ProcessResponse {
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
	// Label objects seen on the image
	tags := image.Label(img.Bytes)
	for k, v := range tags {
		img.Meta[k] = v
	}
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

func S5(workerCnt, imageCnt int) *Stats {
	return Process(workerCnt, imageCnt, ProcessImageS5)
}
