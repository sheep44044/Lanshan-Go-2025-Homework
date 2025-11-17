package main

import (
	"awesomeProject1/homework05/workerpool"
	"fmt"
	"sync/atomic"
)

func main() {
	pool := workerpool.New(workerpool.DefaultConfig())
	defer pool.Close()

	var count int32
	totalTasks := 1000

	// 提交1000个计数任务
	for i := 0; i < totalTasks; i++ {
		pool.SubmitFunc(func() {
			atomic.AddInt32(&count, 1)
		})
	}

	// 等待所有任务完成
	pool.Close()

	fmt.Printf("最终计数: %d (期望: %d)\n", atomic.LoadInt32(&count), totalTasks)
}
