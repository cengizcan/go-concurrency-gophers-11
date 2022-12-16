package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cengizcan/go-concurrency-patterns/steps"
)

const FILE_COUNT = 50

func main() {
	fmt.Println("Go concurrency patterns")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1000))
	defer cancel()
	imageCnt := []int{10, 20, 25, 5}
	fmt.Printf("%v\n", steps.S7(ctx, 3, imageCnt))
}
