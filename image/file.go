package image

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	origSize          = 3000
	origSizeVariation = 1000
)

type MetaItem struct {
	Key, Val string
}

type FakeImage struct {
	Name  string
	Id    int
	Bytes []byte
	Meta  map[string]string
}

func ReadImage(name string) *FakeImage {
	l := rand.Intn(origSizeVariation) + origSize
	time.Sleep(time.Duration(l/10) * time.Microsecond)
	bytes := make([]byte, l)
	rand.Read(bytes)
	return &FakeImage{
		Name:  name,
		Bytes: bytes,
		Meta:  map[string]string{"OrigSize": strconv.Itoa(l)},
	}
}

func ReadSeq(cnt int) []*FakeImage {
	files := make([]*FakeImage, cnt)
	for i := 0; i < cnt; i++ {
		files[i] = ReadImage(fmt.Sprintf("IMG_%d", i))
	}
	return files
}

func WriteImage(img *FakeImage) *FakeImage {
	time.Sleep(time.Duration(len(img.Bytes)/10) * time.Microsecond)
	return img
}

func ScanImages(cnt int) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < cnt; i++ {
			out <- fmt.Sprintf("image_%d.jpg", i)
			//time.Sleep(time.Second)
		}
		close(out)
	}()

	return out
}
