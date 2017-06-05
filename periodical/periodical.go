package periodical

// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"sync"
	"time"

	"github.com/corpix/scheduler/errors"
	"github.com/corpix/scheduler/executor"
	"github.com/corpix/scheduler/work"
)

type Periodical struct {
	*sync.Mutex

	config   Config
	executor executor.Executor
	queue    chan *work.Work
	tasks    map[*work.Work]time.Time
	done     chan struct{}
}

func (p *Periodical) run() {
	go p.plan()
	go p.execute()
}

func (p *Periodical) plan() {
	for {
		p.Lock()
		now := time.Now()
		for work, ts := range p.tasks {
			timeshift := ts.Add(work.Schedule.(*Schedule).Every)
			if timeshift.Equal(now) || timeshift.Before(now) {
				p.tasks[work] = now
				p.queue <- work
			}
		}
		p.Unlock()

		select {
		case <-time.After(p.config.Tick):
		case <-p.done:
			return
		}
	}
}

func (p *Periodical) execute() {
	for {
		select {
		case work, ok := <-p.queue:
			if ok {
				p.executor.Execute(work.Fn)
			}
		case <-p.done:
			return
		}
	}
}

func (p Periodical) Unschedule(w *work.Work) {
	p.Lock()
	defer p.Unlock()
	p.unschedule(w)
}

func (p Periodical) unschedule(w *work.Work) {
	delete(p.tasks, w)
}

func (p *Periodical) Schedule(w *work.Work) error {
	_, ok := w.Schedule.(*Schedule)
	if !ok {
		return errors.NewErrUnknownSchedule(
			&Schedule{},
			w.Schedule,
		)
	}

	p.Lock()
	defer p.Unlock()

	return p.schedule(w)
}

func (p *Periodical) schedule(w *work.Work) error {
	_, ok := p.tasks[w]
	if ok {
		return errors.NewErrAlreadyScheduled(w)
	}

	p.tasks[w] = time.
		Now().
		Add(-1 * w.Schedule.(*Schedule).Every)

	return nil
}

func (p *Periodical) Close() {
	p.Lock()
	defer p.Unlock()

	p.tasks = map[*work.Work]time.Time{}
	close(p.done)
	close(p.queue)
}

func New(executor executor.Executor, config Config) (*Periodical, error) {
	p := &Periodical{
		Mutex: &sync.Mutex{},

		config:   config,
		executor: executor,
		queue:    make(chan *work.Work, config.BacklogSize),
		tasks:    make(map[*work.Work]time.Time),
		done:     make(chan struct{}),
	}
	p.run()
	return p, nil
}
