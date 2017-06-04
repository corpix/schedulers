scheduler
---------

Simple scheduler with extendable scheduling algorithms.

## Example

``` go
package main

import (
	"fmt"
	"time"

	"github.com/corpix/scheduler"
	"github.com/corpix/scheduler/periodical"
	"github.com/corpix/scheduler/work"
)

func main() {
	s, err := scheduler.NewFromConfig(
		func(fn func()) { fn() },
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
```

Now run:

``` console
$ go run ./example/simple/simple.go
I am running 2017-06-04 20:32:51.293647757 +0000 UTC
Me running too 2017-06-04 20:32:51.293740099 +0000 UTC
I am running 2017-06-04 20:32:56.294000904 +0000 UTC
I am running 2017-06-04 20:33:01.29451009 +0000 UTC
Me running too 2017-06-04 20:33:01.2945496 +0000 UTC
I am running 2017-06-04 20:33:06.294949678 +0000 UTC
Me running too 2017-06-04 20:33:11.295449691 +0000 UTC
I am running 2017-06-04 20:33:11.295468104 +0000 UTC
...
```
