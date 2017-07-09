package periodical

import (
	"sync"
	"testing"
	"time"

	"github.com/corpix/scheduler/executor"
	"github.com/corpix/scheduler/executor/inplace"
	"github.com/corpix/scheduler/task"
	"github.com/stretchr/testify/assert"
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

	task5emited := false
	task5 := task.New(
		&Schedule{Every: 300 * time.Millisecond},
		func() {
			if !task5emited {
				task5emited = true
				w.Done()
			}
		},
	)
	err = s.Schedule(task5)
	if err != nil {
		panic(err)
	}

	task10emited := false
	task10 := task.New(
		&Schedule{Every: 300 * time.Millisecond},
		func() {
			if !task10emited {
				task10emited = true
				w.Done()
			}
		},
	)
	err = s.Schedule(task10)
	if err != nil {
		panic(err)
	}

	w.Wait()

	assert.Equal(t, true, task5emited)
	assert.Equal(t, true, task10emited)
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
	w.Add(2)

	task5emited := 0
	task5 := task.New(
		&Schedule{Every: 1 * time.Microsecond},
		func() {
			task5emited++
			w.Done()
		},
	)
	err = s.Schedule(task5)
	if err != nil {
		t.Error(err)
		return
	}

	task10emited := 0
	task10 := task.New(
		&Schedule{Every: 1 * time.Microsecond},
		func() {
			task10emited++
			w.Done()
		},
	)
	err = s.Schedule(task10)
	if err != nil {
		t.Error(err)
		return
	}

	w.Wait()

	s.Unschedule(task5)
	s.Unschedule(task10)

	time.Sleep(200 * time.Millisecond)

	w.Add(2)
	err = s.Schedule(task5)
	if err != nil {
		t.Error(err)
		return
	}
	err = s.Schedule(task10)
	if err != nil {
		t.Error(err)
		return
	}

	w.Wait()

	s.Unschedule(task5)
	s.Unschedule(task10)

	assert.Equal(t, 2, task5emited)
	assert.Equal(t, 2, task10emited)
}
