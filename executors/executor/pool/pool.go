package pool

import (
	"context"

	"github.com/corpix/pool"

	"github.com/corpix/schedulers/work"
)

type Pool struct {
	pool *pool.Pool
}

func (p *Pool) Execute(fn work.Work) {
	p.pool.Feed <- pool.NewWork(
		context.Background(),
		func(c context.Context) {
			select {
			case <-c.Done():
			default:
				fn()
			}
		},
	)
}

func NewFromConfig(c Config) (*Pool, error) {
	return &Pool{
		pool: pool.New(
			c.Workers,
			c.QueueSize,
		),
	}, nil
}
