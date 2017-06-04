package errors

// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/corpix/scheduler/work"
)

type ErrUnknownSchedule struct {
	want interface{}
	got  interface{}
}

func (e *ErrUnknownSchedule) Error() string {
	return fmt.Sprintf(
		"Unknown schedule type, want '%T' got '%T'",
		e.want,
		e.got,
	)
}
func NewErrUnknownSchedule(want, got interface{}) error {
	return &ErrUnknownSchedule{want, got}
}

//

type ErrAlreadyScheduled struct {
	work *work.Work
}

func (e *ErrAlreadyScheduled) Error() string {
	return fmt.Sprintf(
		"Work '%s' already scheduled",
		spew.Sdump(e.work),
	)
}
func NewErrAlreadyScheduled(w *work.Work) error {
	return &ErrAlreadyScheduled{w}
}

//
