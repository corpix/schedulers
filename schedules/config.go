package schedules

import (
	"github.com/corpix/schedulers/schedules/schedule/periodical"
)

type Config struct {
	Type       string
	Periodical periodical.Schedule
}
