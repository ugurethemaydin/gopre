package gopre

import (
	"github.com/davecgh/go-spew/spew"
	"os"
)

func pre(x interface{}, y ...interface{}) {
	//rw()
	spew.Dump(x)
	die()
}

func die() {
	os.Exit(1)
}
