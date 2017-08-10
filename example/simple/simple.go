package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/corpix/schedulers"
	"github.com/corpix/schedulers/executors"
	"github.com/corpix/schedulers/executors/executor/inplace"
	"github.com/corpix/schedulers/scheduler/periodical"
	schedule "github.com/corpix/schedulers/schedules/schedule/periodical"
	"github.com/corpix/schedulers/task"
)

func main() {
	e, err := executors.NewFromConfig(
		executors.Config{
			Type:    executors.InplaceExecutorType,
			Inplace: inplace.Config{},
		},
	)
	if err != nil {
		panic(err)
	}

	s, err := schedulers.NewFromConfig(
		schedulers.Config{
			Type: schedulers.PeriodicalSchedulerType,
			Periodical: periodical.Config{
				Tick:      1 * time.Second,
				QueueSize: 5,
			},
		},
		e,
	)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var (
		wg = &sync.WaitGroup{}

		task1Counter = 4
		task2Counter = 4

		task1 *task.Task
		task2 *task.Task
	)

	task1 = task.New(
		schedule.Schedule{Every: 1 * time.Second},
		func() {
			if task1Counter == 0 {
				s.Unschedule(task1)
				wg.Done()
				return
			}

			fmt.Println("I am running", time.Now())
			task1Counter--
		},
	)
	wg.Add(1)

	task2 = task.New(
		schedule.Schedule{Every: 3 * time.Second},
		func() {
			if task2Counter == 0 {
				s.Unschedule(task2)
				wg.Done()
				return
			}

			fmt.Println("Me running too", time.Now())
			task2Counter--
		},
	)
	wg.Add(1)

	err = s.Schedule(task1)
	if err != nil {
		panic(err)
	}

	err = s.Schedule(task2)
	if err != nil {
		panic(err)
	}

	wg.Wait()
	fmt.Println("Done")
}
