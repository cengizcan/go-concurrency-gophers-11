package image

import (
	"math/rand"
	"time"
)

const (
	reSize          = 2000
	reSizeVariation = 750
)

func Resize(img *FakeImage) *FakeImage {
	l := rand.Intn(reSizeVariation) + reSize
	time.Sleep(time.Duration(l/100) * time.Microsecond)
	bytes := make([]byte, l)
	rand.Read(bytes)
	return &FakeImage{
		Name:  img.Name,
		Id:    img.Id,
		Bytes: bytes,
		Meta:  img.Meta,
	}
}
