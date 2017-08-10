package periodical

import (
	"sync"
	"testing"
	"time"

	"github.com/corpix/schedulers/executors"
	"github.com/corpix/schedulers/executors/executor/inplace"
	schedule "github.com/corpix/schedulers/schedules/schedule/periodical"
	"github.com/corpix/schedulers/task"
)

func TestSchedule(t *testing.T) {
	e, err := executors.NewFromConfig(
		executors.Config{
			Type:    "inplace",
			Inplace: inplace.Config{},
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	s, err := NewFromConfig(
		Config{
			Tick:      100 * time.Millisecond,
			QueueSize: 5,
		},
		e,
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
			schedule.Schedule{Every: 300 * time.Millisecond},
			func() { w.Done() },
		),
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.Schedule(
		task.New(
			schedule.Schedule{Every: 300 * time.Millisecond},
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
	e, err := executors.NewFromConfig(
		executors.Config{
			Type:    "inplace",
			Inplace: inplace.Config{},
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	s, err := NewFromConfig(
		Config{
			Tick:      100 * time.Millisecond,
			QueueSize: 5,
		},
		e,
	)
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()

	w := &sync.WaitGroup{}

	task1 := task.New(
		schedule.Schedule{Every: 1 * time.Microsecond},
		func() { w.Done() },
	)
	err = s.Schedule(task1)
	if err != nil {
		t.Error(err)
		return
	}

	task2 := task.New(
		schedule.Schedule{Every: 1 * time.Microsecond},
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
