package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/corpix/pool"
)

func main() {
	p := pool.New(10, 10)
	defer p.Close()

	w := &sync.WaitGroup{}

	tasks := 10
	sleep := 1 * time.Second

	for n := 0; n < tasks; n++ {
		w.Add(1)
		p.Feed <- pool.NewWork(
			context.Background(),
			func(n int) pool.Executor {
				return func(ctx context.Context) {
					defer w.Done()
					select {
					case <-ctx.Done():
					default:
						time.Sleep(sleep)
						fmt.Printf("Finished work '%d'\n", n)
					}
				}
			}(n),
		)
	}

	w.Wait()
}
