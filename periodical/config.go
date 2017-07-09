package periodical

import (
	"time"
)

type Config struct {
	Tick        time.Duration
	BacklogSize int8
}
