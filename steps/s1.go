package steps

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Basic generator
func generator() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("iteration %d", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// A half baked real-world file scanner
func scan(dir string) <-chan string {
	out := make(chan string)

	go func() {
		// Alternative: filepath.Walk
		files, err := ioutil.ReadDir(dir)
		// TODO: Handle errors gracefully
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			path := fmt.Sprintf("%s/%s", dir, f.Name())
			if f.IsDir() {
				// walk through the child directory
				for p := range scan(path) {
					out <- p
				}
			} else if strings.HasSuffix(f.Name(), ".jpg") { // TODO: lookup -> jpg, jpeg, uppercase...
				out <- path
			}
		}
		close(out)
	}()

	return out
}

func fakeScanner(cnt int) <-chan string {
	out := make(chan string)
	for i := 0; i < cnt; i++ {
		out <- fmt.Sprintf("image_%d.jpg", i)
	}
	return out
}
