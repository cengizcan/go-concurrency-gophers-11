package steps

import (
	"context"
	"sync"

	"github.com/cengizcan/go-concurrency-patterns/image"
	"github.com/cengizcan/go-concurrency-patterns/numerator"
)

// Generic worker pool
type WorkerPool[T any] struct {
	workerCnt int
	Tasks     chan Executable[T]
	Results   chan T
}

func New[T any](cnt int) *WorkerPool[T] {
	return &WorkerPool[T]{
		workerCnt: cnt,
		Tasks:     make(chan Executable[T]),
		Results:   make(chan T),
	}
}

func (wp *WorkerPool[T]) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workerCnt; i++ {
		wg.Add(1)
		go worker(ctx, wp.Tasks, wp.Results, &wg)
	}
	wg.Wait()
	close(wp.Results)
}

// Worker
func worker[T any](ctx context.Context, tasks <-chan Executable[T], results chan<- T, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				return
			}
			results <- task.Execute()
		case <-ctx.Done():
			//fmt.Printf("Context done received: %v\n", ctx.Err())
			return
		}
	}
}

type Executable[T any] interface {
	Execute() T
}

// Our specific task
type Task struct {
	path string
	fn   ProcessImage
	num  numerator.Sequential
}

func (t *Task) Execute() *ProcessResponse {
	return t.fn(t.path, t.num)
}

func NewTask(path string, fn ProcessImage, num numerator.Sequential) *Task {
	return &Task{
		path: path,
		fn:   fn,
		num:  num,
	}
}

// Run
func S6(ctx context.Context, workerCnt, imageCnt int) *Stats {
	num := numerator.NewV2()

	wp := New[*ProcessResponse](workerCnt)
	go wp.Run(ctx)

	go func() {
		for p := range image.ScanImages(imageCnt) {
			wp.Tasks <- NewTask(p, ProcessImageS5, num)
		}
		close(wp.Tasks)
	}()

	return StatGen(wp.Results)
}

// Worker with done channel
func worker2[T any](ctx context.Context, tasks <-chan Executable[T], results chan<- T, done chan<- bool) {
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				done <- true
				return
			}
			results <- task.Execute()
		case <-ctx.Done():
			//fmt.Printf("Context done received: %v\n", ctx.Err())
			return
		}
	}
}

// Run with done channel
func (wp *WorkerPool[T]) Run2(ctx context.Context) {
	done := make(chan bool, wp.workerCnt)

	for i := 0; i < wp.workerCnt; i++ {

		go worker2(ctx, wp.Tasks, wp.Results, done)
	}

	for i := 0; i < wp.workerCnt; i++ {
		<-done
	}
	close(wp.Results)
}

// Concurrent implementation
func S6Concurrent(cnt int) {
	wg := new(sync.WaitGroup)
	for p := range image.ScanImages(cnt) {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			ProcessImageS5(path, numerator.NewV2())
		}(p)
	}
	wg.Wait()
}
