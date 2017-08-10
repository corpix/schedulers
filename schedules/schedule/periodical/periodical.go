package periodical

import (
	"time"
)

type Schedule struct {
	Every time.Duration
}

func New(e time.Duration) Schedule {
	return Schedule{Every: e}
}
