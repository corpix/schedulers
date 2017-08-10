package periodical

import (
	"sync"
	"time"

	"github.com/corpix/schedulers/errors"
	"github.com/corpix/schedulers/executors/executor"
	"github.com/corpix/schedulers/task"
)

type Periodical struct {
	*sync.Mutex
	config   Config
	executor executor.Executor
	queue    chan *task.Task
	tasks    map[*task.Task]time.Time
	done     chan struct{}
}

func (p *Periodical) run() {
	go p.plan()
	go p.execute()
}

func (p *Periodical) plan() {
	for {
		var (
			now = time.Now()
		)

		p.Lock()
		for t, ts := range p.tasks {
			timeshift := ts.Add(t.Schedule.(*Schedule).Every)
			if timeshift.Equal(now) || timeshift.Before(now) {
				p.tasks[t] = now
				select {
				case p.queue <- t:
				}
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
		case t, ok := <-p.queue:
			if ok {
				p.executor.Execute(t.Fn)
			}
		case <-p.done:
			return
		}
	}
}

func (p *Periodical) Schedule(w *task.Task) error {
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

func (p *Periodical) schedule(w *task.Task) error {
	_, ok := p.tasks[w]
	if ok {
		return errors.NewErrAlreadyScheduled(w)
	}

	p.tasks[w] = time.
		Now().
		Add(-1 * w.Schedule.(*Schedule).Every)

	return nil
}

func (p *Periodical) Unschedule(w *task.Task) {
	p.Lock()
	defer p.Unlock()
	p.unschedule(w)
}

func (p *Periodical) unschedule(w *task.Task) {
	delete(p.tasks, w)
}

func (p *Periodical) Close() {
	p.Lock()
	defer p.Unlock()

	p.tasks = map[*task.Task]time.Time{}
	close(p.queue)
	close(p.done)
}

func NewFromConfig(c Config, e executor.Executor) (*Periodical, error) {
	p := &Periodical{
		Mutex:    &sync.Mutex{},
		config:   c,
		executor: e,
		queue:    make(chan *task.Task, c.QueueSize),
		tasks:    make(map[*task.Task]time.Time),
		done:     make(chan struct{}),
	}
	p.run()
	return p, nil
}
