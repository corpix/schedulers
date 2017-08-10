package schedules

import (
	"strings"

	"github.com/corpix/schedulers/schedules/schedule"
)

const (
	PeriodicalScheduleType = "periodical"
)

func NewFromConfig(c Config) (schedule.Schedule, error) {
	var (
		t = strings.ToLower(c.Type)
	)

	switch t {
	case PeriodicalScheduleType:
		return c.Periodical, nil
	default:
		return nil, NewErrUnknownScheduleType(c.Type)
	}
}
