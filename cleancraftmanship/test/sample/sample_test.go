package sample

import (
	"testing"
)

func TestCanCreateStack(t * testing.T) {
  s := StackTest.CanCreateStack()

  if s.IsEmpty() {
    t.Error("FAIL")
  }
}
