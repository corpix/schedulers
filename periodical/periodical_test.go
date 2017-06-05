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
