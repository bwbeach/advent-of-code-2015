package main

import (
	"testing"
)

func Test_computeFloor(t *testing.T) {
	data := []struct {
		text            string
		expected        int
		expectedMessage string
	}{
		{"(())", 0, ""},
		{"()()", 0, ""},
		{"(()(()(", 3, ""},
		{"))(((((", 3, ""},
		{")())())", -3, ""},
		{"(X)", 0, "bad character"},
	}
	for _, d := range data {
		floor, err := computeFloor(d.text)
		if floor != d.expected {
			t.Errorf("for '%s' expected %d but got %d", d.text, d.expected, floor)
		}
		actualMsg := ""
		if err != nil {
			actualMsg = err.Error()
		}
		if actualMsg != d.expectedMessage {
			t.Errorf("for '%s' expected error '%s' but got '%s'", d.text, d.expectedMessage, actualMsg)
		}
	}
}