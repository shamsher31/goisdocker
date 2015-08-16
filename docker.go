// Package docker check if the process is running inside a Docker container
package docker // import "github.com/shamsher31/goisdocker"

import (
	"os"
)

func Is() bool {
	if _, err := os.Stat("/.dockerinit"); os.IsNotExist(err) {
		return false
	}
	return true
}
