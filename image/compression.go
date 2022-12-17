package image

import (
	"fmt"
	"math/rand"
	"time"
)

type CompResult struct {
	SvcName string
	Output  []byte
}

type Compressor func(src []byte) *CompResult

var (
	services = []Compressor{
		externalCompress("Svc1"),
		externalCompress("Svc2"),
		externalCompress("Svc3"),
		externalCompress("Svc4"),
	}
)

func externalCompress(svc string) Compressor {
	return func(src []byte) *CompResult {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		bytes := make([]byte, len(src)/5)
		rand.Read(bytes)
		return &CompResult{
			SvcName: svc,
			Output:  bytes,
		}
	}
}

func Compress(src []byte) *CompResult {
	result := make(chan *CompResult)

	for _, comp := range services {
		go func(compress Compressor) {
			res := compress(src)

			select {
			case result <- res:
			default:
			}
		}(comp)
	}
	// Instead of returning Result chan, returns first result
	return <-result
}

// Compress with quit channel
func compressWithDone(quit chan bool, compress Compressor, src []byte, result chan<- *CompResult) {
	response := make(chan *CompResult)
	//err := make(chan error)
	go func() {
		response <- compress(src)
		//response, err <- compress(src)
	}()

	select {
	case result <- <-response:
		// equal to:
		// a := <-response
		//  resutl <- a
	case <-quit:
		return
	}
}

func Compress2(src []byte) *CompResult {
	quit := make(chan bool, 1)
	result := make(chan *CompResult)

	for _, comp := range services {
		go compressWithDone(quit, comp, src, result)
	}

	firstReturned := <-result
	close(quit)
	fmt.Println(firstReturned.SvcName)
	return firstReturned
}
