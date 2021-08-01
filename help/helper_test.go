package main

import (
	"testing"

	"github.com/13sai/gohelper"
)

func TestHelper(t *testing.T) {
	t.Log(gohelper.StrCombine("hi,", "sai!"))
}
