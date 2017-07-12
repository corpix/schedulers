package pool

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoolParallel(t *testing.T) {
	tasks := 10
	workers := 10
	queue := 10
	sleep := 100 * time.Millisecond

	p := New(workers, queue)
	defer p.Close()

	w := &sync.WaitGroup{}
	w.Add(tasks)

	started := time.Now()
	for n := 0; n < tasks; n++ {
		p.Feed <- NewWork(
			context.Background(),
			func(ctx context.Context) {
				select {
				case <-ctx.Done():
				case <-time.After(sleep):
				}
				w.Done()
			},
		)
	}
	w.Wait()
	finished := time.Now()

	assert.False(
		t,
		started.Add(sleep*time.Duration(tasks)).Before(finished),
	)
}

func TestPoolContextCancel(t *testing.T) {
	tasks := 5
	workers := 5
	queue := 0

	p := New(workers, queue)
	defer p.Close()

	w := &sync.WaitGroup{}
	w.Add(tasks)

	cancels := make(chan int, tasks*2)
	defer close(cancels)

	for n := 0; n < tasks; n++ {
		ctx, cancel := context.WithCancel(
			context.Background(),
		)
		cancel()

		p.Feed <- NewWork(
			ctx,
			func(ctx context.Context) {
				select {
				case <-ctx.Done():
					cancels <- 1
				}
				w.Done()
			},
		)
	}

	canceled := 0
	go func() {
		for c := range cancels {
			canceled += c
			if canceled == tasks {
				break
			}
		}
		w.Done()
	}()
	w.Add(1)

	w.Wait()

	assert.Equal(t, tasks, canceled)
}

func TestPoolWithResult(t *testing.T) {
	tasks := 10
	workers := 10
	queue := 10
	results := make(chan *Result)

	errors := tasks / 2
	successes := tasks - errors

	p := New(workers, queue)
	defer p.Close()

	for n := 1; n <= tasks; n++ {
		p.Feed <- NewWorkWithResult(
			context.Background(),
			results,
			func(n int) ExecutorWithResult {
				return func(ctx context.Context) (interface{}, error) {
					if n <= errors {
						return nil, fmt.Errorf("some error")
					}
					return n, nil
				}
			}(n),
		)
	}

	achievedErrors := 0
	achievedSuccesses := 0
	for tasks > 0 {
		result := <-results
		if result.Err != nil {
			achievedErrors++
		}
		if result.Value != nil {
			achievedSuccesses++
		}
		tasks--
	}

	assert.Equal(t, achievedErrors, errors)
	assert.Equal(t, achievedSuccesses, successes)

	assert.Equal(t, 0, tasks)
}
