package schedulers

import (
	"github.com/corpix/schedulers/scheduler/periodical"
)

type Config struct {
	Type       string
	Periodical periodical.Config
}
