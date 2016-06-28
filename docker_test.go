package docker

import "testing"

func TestIs(t *testing.T) {
	if Is() == false {
		t.Fatal("Your application is not running inside docker conatiner")
	}
}
