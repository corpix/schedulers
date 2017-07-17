package periodical

import (
	"sync"
	"testing"
	"time"

	"github.com/corpix/scheduler/executor"
	"github.com/corpix/scheduler/executor/inplace"
	"github.com/corpix/scheduler/task"
)

func TestSchedule(t *testing.T) {
	e, err := executor.NewFromConfig(inplace.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	s, err := New(
		e,
		Config{
			Tick:        100 * time.Millisecond,
			BacklogSize: 5,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()

	w := &sync.WaitGroup{}
	w.Add(2)

	err = s.Schedule(
		task.New(
			&Schedule{Every: 300 * time.Millisecond},
			func() { w.Done() },
		),
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.Schedule(
		task.New(
			&Schedule{Every: 300 * time.Millisecond},
			func() { w.Done() },
		),
	)
	if err != nil {
		t.Error(err)
		return
	}

	w.Wait()
}

func TestScheduleUnschedule(t *testing.T) {
	e, err := executor.NewFromConfig(inplace.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	s, err := New(
		e,
		Config{
			Tick:        100 * time.Millisecond,
			BacklogSize: 5,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()

	w := &sync.WaitGroup{}

	task1 := task.New(
		&Schedule{Every: 1 * time.Microsecond},
		func() { w.Done() },
	)
	err = s.Schedule(task1)
	if err != nil {
		t.Error(err)
		return
	}

	task2 := task.New(
		&Schedule{Every: 1 * time.Microsecond},
		func() { w.Done() },
	)
	err = s.Schedule(task2)
	if err != nil {
		t.Error(err)
		return
	}

	s.Unschedule(task1)
	s.Unschedule(task2)

	w.Wait()
}
