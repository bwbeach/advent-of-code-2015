package main

import (
	"testing"
)

func Test_computeFloor(t *testing.T) {
	floor, err := computeFloor("((()())")
	if err != nil {
		t.Error("did not expect error", err)
	}
	if floor != 1 {
		t.Error("expected floor 1", floor)
	}
}
