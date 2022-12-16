package numerator

import (
	"sync"
	"sync/atomic"
)

type Sequential interface {
	Next() int
}
type sequenceV1 struct {
	current int
}

func (s *sequenceV1) Next() int {
	s.current++
	return s.current
}

func NewV1() Sequential {
	return &sequenceV1{
		current: 0,
	}
}

/*
* MUTEX
 */
type sequenceV2 struct {
	mu      sync.Mutex
	current int
}

func (s *sequenceV2) Next() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.current++
	return s.current
}

func NewV2() Sequential {
	return &sequenceV2{
		current: 0,
	}
}

/*
* ATOMIC
 */
type sequenceV3 struct {
	current int32
}

func (s *sequenceV3) Next() int {
	atomic.AddInt32(&s.current, 1)
	return int(s.current)
}
func NewV3() Sequential {
	return &sequenceV3{}
}
