package image

import (
	"fmt"
	"math/rand"
	"time"
)

type LabelResult struct {
	AlgoName string
	Value    string
}
type LabelFn func(src []byte, timeout <-chan struct{}, results chan *LabelResult)

var (
	algos = []LabelFn{
		externalAlgo("AlgoA"),
		externalAlgo("AlgoB"),
		externalAlgo("AlgoC"),
		externalAlgo("AlgoD"),
	}
)

func fakeAlgo(svc string) *LabelResult {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	return &LabelResult{
		AlgoName: svc,
		Value:    fmt.Sprintf("Value of %s", svc),
	}
}
func externalAlgo(svc string) LabelFn {
	return func(src []byte, timeout <-chan struct{}, results chan *LabelResult) {
		results <- fakeAlgo(svc)
	}
}

func Label(src []byte) map[string]string {
	timeout := make(chan struct{}, 1)
	results := make(chan *LabelResult, len(algos))

	go func() {
		time.Sleep(15 * time.Millisecond)
		timeout <- struct{}{}
	}()

	for _, t := range algos {
		go t(src, timeout, results)
	}

	resMap := make(map[string]string)
	for {
		select {
		case r := <-results:
			resMap[r.AlgoName] = r.Value

			if len(resMap) == len(algos) {
				return resMap
			}
		case <-timeout:
			return resMap
		}
	}

}
