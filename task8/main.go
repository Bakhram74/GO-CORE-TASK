package main

import (
	"fmt"
	"sync/atomic"
)

type CustomWaitGroup struct {
	counter int32
	done    chan struct{}
}

func NewWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		counter: 0,
		done:    make(chan struct{}),
	}
}
func (wg *CustomWaitGroup) Add(n int) {
	count := atomic.AddInt32(&wg.counter, int32(n))
	if count < 0 {
		panic("negative number has provided")
	}
	if count == 0 {
		wg.done <- struct{}{}
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}

func main() {
	wg := NewWaitGroup()

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Got: ", i)
		}()
	}
	wg.Wait()
}
