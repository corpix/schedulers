package inplace

import (
	"github.com/corpix/schedulers/work"
)

type Inplace struct{}

func (e *Inplace) Execute(fn work.Work) {
	fn()
}

func NewFromConfig(c Config) (*Inplace, error) {
	return &Inplace{}, nil
}
