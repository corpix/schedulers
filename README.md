schedulers
----------

[![Build Status](https://travis-ci.org/corpix/schedulers.svg?branch=master)](https://travis-ci.org/corpix/schedulers)

Simple scheduler with extendable scheduling algorithms.

Supported algorithms:

- `periodical` schedules tasks periodically, e.g. "every 5 minutes"

## Example

``` console
$ go run ./example/simple/simple.go
I am running 2017-08-10 23:22:04.137815031 +0000 UTC
Me running too 2017-08-10 23:22:05.137849629 +0000 UTC
I am running 2017-08-10 23:22:05.137869246 +0000 UTC
I am running 2017-08-10 23:22:06.137976179 +0000 UTC
I am running 2017-08-10 23:22:07.138115222 +0000 UTC
Me running too 2017-08-10 23:22:08.138184934 +0000 UTC
Me running too 2017-08-10 23:22:11.138478238 +0000 UTC
Me running too 2017-08-10 23:22:14.13880726 +0000 UTC
Done
```
