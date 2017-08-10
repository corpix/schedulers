package periodical

import (
	"time"
)

type Config struct {
	Tick      time.Duration
	QueueSize int8
}
