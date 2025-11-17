package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32 = 0
var totalIncrements int = 1000

type Task struct {
	Runnable func()
}

func main() {
	var wg sync.WaitGroup
	channels := make([]chan Task, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan Task, 10)
		wg.Add(1)
		go func(workID int) {
			defer wg.Done()
			for t := range channels[i] {
				t.Runnable()
			}
		}(i)
	}

	for i := range 1000 {
		workerIndex := i % 10
		t1 := Task{
			Runnable: func() {
				atomic.AddInt32(&count, 1)
			},
		}
		channels[workerIndex] <- t1
	}

	for i := 0; i < 10; i++ {
		close(channels[i])
	}
	wg.Wait()

	fmt.Printf("最终结果: %d (期望: %d)\n", atomic.LoadInt32(&count), totalIncrements)
}
