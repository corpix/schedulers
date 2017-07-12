package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/corpix/pool"
)

func main() {
	p := pool.New(10, 10)
	defer p.Close()

	w := &sync.WaitGroup{}

	tasks := 10
	results := make(chan *pool.Result)
	defer close(results)

	for n := 0; n < tasks; n++ {
		w.Add(1)
		p.Feed <- pool.NewWorkWithResult(
			context.Background(),
			results,
			func(n int) pool.ExecutorWithResult {
				return func(ctx context.Context) (interface{}, error) {
					select {
					case <-ctx.Done():
						return nil, ctx.Err()
					default:
						fmt.Printf("Finished work '%d'\n", n)
					}
					return n, nil
				}
			}(n),
		)
	}

	go func() {
		// Releasing one worker per iteration
		// when using not buffered channels.
		for result := range results {
			fmt.Printf(
				"Received result '%d'\n",
				result.Value.(int),
			)
			w.Done()
		}
	}()

	w.Wait()
}
