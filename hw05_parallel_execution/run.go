package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {

	var (
		wg         *sync.WaitGroup
		currentErr int64
	)
	chanTask := make(chan Task)

	defer func() {
		close(chanTask)
		wg.Wait()
	}()

	wg.Add(n)

	for i := 0; i < n; i++ {

		go func() {
			for tsk := range chanTask {
				if err := tsk(); err != nil {
					atomic.AddInt64(&currentErr, 1)
				}
			}
			wg.Done()
		}()

	}

	for _, task := range tasks {
		if atomic.LoadInt64(&currentErr) >= int64(m) {
			return ErrErrorsLimitExceeded
		}
		chanTask <- task
	}

	return nil
}
