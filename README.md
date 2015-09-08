# goisdocker

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goisdocker)
[![Build Status](https://travis-ci.org/shamsher31/goisdocker.svg)](https://travis-ci.org/shamsher31/goisdocker)
[![GitHub release](http://img.shields.io/github/release/shamsher31/goisdocker.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)

Check if the process is running inside a Docker container

### How to install
```go
go get github.com/shamsher31/goisdocker
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goisdocker"
)

func main() {
    if(docker.Is()) {
    fmt.Println("Running inside a Docker container")
  }
}
```
###Aliasing Imports
If you already have package name ```docker``` you can use following.
```go
package main

import (
	"fmt"
	pckDocker "github.com/shamsher31/goisdocker"
)

func main() {
    if(pckDocker.Is()) {
    fmt.Println("Running inside a Docker container")
  }
}
```

### Related
[goisvdo](https://github.com/shamsher31/goisvdo)<br>
[goistext](https://github.com/shamsher31/goistext)<br>
[goisimg](https://github.com/shamsher31/goisimg)<br>

### Why
This package is inspired by [is-docker](https://www.npmjs.com/package/is-docker) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
