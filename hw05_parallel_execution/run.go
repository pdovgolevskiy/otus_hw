package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) < n {
		n = len(tasks)
	}
	var errorCount int32
	taskChannel := make(chan Task)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for currentTask := range taskChannel {
				if err := currentTask(); err != nil {
					atomic.AddInt32(&errorCount, 1)
				}
			}
		}()
	}
	//
	for _, task := range tasks {
		if atomic.LoadInt32(&errorCount) >= int32(m) {
			break
		}
		taskChannel <- task
	}
	close(taskChannel)
	wg.Wait()
	if atomic.LoadInt32(&errorCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
