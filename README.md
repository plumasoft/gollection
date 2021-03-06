# gollection
[![GoDoc](https://godoc.org/gopkg.in/azihsoyn/gollection.v1?status.svg)](https://godoc.org/gopkg.in/azihsoyn/gollection.v1)
[![Coverage Status](https://coveralls.io/repos/github/azihsoyn/gollection/badge.svg?branch=master)](https://coveralls.io/github/azihsoyn/gollection?branch=master)
[![CircleCI](https://circleci.com/gh/azihsoyn/gollection.svg?style=svg)](https://circleci.com/gh/azihsoyn/gollection)

gollection is golang collection util library (inspired by [kotlin collection](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.collections/index.html))

# supported functions
- distinct
- distinctBy
- filter
- flatMap
- flatten
- fold
- map
- reduce
- skip ([java8](https://docs.oracle.com/javase/8/docs/api/java/util/stream/Stream.html#skip-long-))
- sortBy
- take

# Feature

- read-only original slice
- stream(goroutine) support

# Installation

`go get github.com/azihsoyn/gollection` or `go get gopkg.in/azihsoyn/gollection.v1`

# Usage
see [examples code](https://github.com/azihsoyn/gollection/tree/master/examples)

or [godoc](https://coveralls.io/github/azihsoyn/gollection?branch=master) [examples](https://godoc.org/github.com/azihsoyn/gollection#pkg-examples)

# LICENSE
MIT
