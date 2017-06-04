package main

import (
	"fmt"
	"time"

	"github.com/corpix/scheduler"
	"github.com/corpix/scheduler/executor"
	"github.com/corpix/scheduler/executor/inplace"
	"github.com/corpix/scheduler/periodical"
	"github.com/corpix/scheduler/work"
)

func main() {
	e, err := executor.NewFromConfig(inplace.Config{})
	if err != nil {
		panic(err)
	}

	s, err := scheduler.NewFromConfig(
		e,
		periodical.Config{
			Tick:        1 * time.Second,
			BacklogSize: 5,
		},
	)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	err = s.Schedule(
		work.New(
			&periodical.Schedule{Every: 5 * time.Second},
			func() {
				fmt.Println("I am running", time.Now())
			},
		),
	)
	if err != nil {
		panic(err)
	}

	err = s.Schedule(
		work.New(
			&periodical.Schedule{Every: 10 * time.Second},
			func() {
				fmt.Println("Me running too", time.Now())
			},
		),
	)
	if err != nil {
		panic(err)
	}

	select {}
}
