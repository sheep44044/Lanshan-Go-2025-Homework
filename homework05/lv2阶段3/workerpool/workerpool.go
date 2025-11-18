package workerpool

import (
	"sync"
	"sync/atomic"
)

var roundRobinCounter int32

type Task struct {
	Runnable func()
}

type Config struct {
	workers       int
	channelbuffer int
}

func DefaultConfig() Config {
	return Config{
		workers:       10,
		channelbuffer: 10,
	}
}

type Workerpool struct {
	workers  int
	channels []chan Task
	wg       sync.WaitGroup
	isClosed bool
	mu       sync.RWMutex
}

func New(config Config) *Workerpool {
	if config.workers <= 0 {
		config.workers = 10
	}
	if config.channelbuffer <= 0 {
		config.channelbuffer = 10
	}
	wp := &Workerpool{
		workers:  config.workers,
		channels: make([]chan Task, config.channelbuffer),
		isClosed: false,
	}
	for i := 0; i < wp.workers; i++ {
		wp.channels[i] = make(chan Task, config.channelbuffer)
		wp.wg.Add(1)
		go wp.Worker(i)
	}
	return wp
}

func (wp *Workerpool) Worker(workerId int) {
	defer wp.wg.Done()
	for task := range wp.channels[workerId] {
		task.Runnable()
	}
}

func (wp *Workerpool) Submit(task Task) bool {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.isClosed {
		return false
	}

	workerIndex := atomic.AddInt32(&roundRobinCounter, 1) % int32(wp.workers)

	wp.channels[workerIndex] <- task
	return true
}

func (wp *Workerpool) SubmitFunc(fn func()) bool {
	return wp.Submit(Task{Runnable: fn})
}

func (wp *Workerpool) Close() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.isClosed {
		return
	}

	wp.isClosed = true
	for i := 0; i < wp.workers; i++ {
		close(wp.channels[i])
	}
	wp.wg.Wait()
}
