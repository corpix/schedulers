schedulers
----------

[![Build Status](https://travis-ci.org/corpix/schedulers.svg?branch=master)](https://travis-ci.org/corpix/schedulers)

Simple scheduler with extendable scheduling algorithms.

Supported algorithms:

- `perodical` schedules tasks periodically, e.g. "every 5 minutes"

## Example

``` console
$ go run ./example/simple/simple.go
I am running 2017-08-10 10:24:17.253772524 +0000 UTC
Me running too 2017-08-10 10:24:18.253845055 +0000 UTC
I am running 2017-08-10 10:24:22.254274021 +0000 UTC
I am running 2017-08-10 10:24:27.254808939 +0000 UTC
Me running too 2017-08-10 10:24:28.254935728 +0000 UTC
I am running 2017-08-10 10:24:32.255383238 +0000 UTC
Me running too 2017-08-10 10:24:38.256078048 +0000 UTC
Me running too 2017-08-10 10:24:48.257107815 +0000 UTC
Done
```
