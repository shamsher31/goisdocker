# goisdocker
Check if the process is running inside a Docker container

# How to install
<pre>
go get github.com/shamsher31/goisdocker
</pre>

# How to use
<pre>
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
</pre>

# Related
[goisvdo](https://github.com/shamsher31/goisvdo)
[goistext](https://github.com/shamsher31/goistext)
[goisimg](https://github.com/shamsher31/goisimg)

# Why
This package is inspired by [is-docker](https://www.npmjs.com/package/is-docker) npm module to check if the process is running inside a Docker container.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
