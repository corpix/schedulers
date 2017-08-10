package schedulers

import (
	"strings"

	"github.com/corpix/schedulers/errors"
	"github.com/corpix/schedulers/executors/executor"
	"github.com/corpix/schedulers/scheduler"
	"github.com/corpix/schedulers/scheduler/periodical"
)

const (
	PeriodicalSchedulerType = "periodical"
)

func NewFromConfig(c Config, e executor.Executor) (scheduler.Scheduler, error) {
	if e == nil {
		return nil, errors.NewErrNilArgument(e)
	}

	switch strings.ToLower(c.Type) {
	case PeriodicalSchedulerType:
		return periodical.NewFromConfig(
			c.Periodical,
			e,
		)
	default:
		return nil, NewErrUnknownSchedulerType(c.Type)
	}
}
