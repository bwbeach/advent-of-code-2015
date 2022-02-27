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

func Test_basementPosition(t *testing.T) {
	pos, err := basementPosition("()())")
	if pos != 5 {
		t.Error("expected pos 5")
	}
	if err != nil {
		t.Error("expected no error")
	}
}

func Test_wrappingPaperNeeded(t *testing.T) {
	data := []struct {
		text            string
		expected        int
		expectedMessage string
	}{
		{"2x3x4", 58, ""},
		{"3x4x2", 58, ""},
		{"3x1", 0, "bad package spec"},
		{"3xAx4", 0, "strconv.Atoi: parsing \"A\": invalid syntax"},
	}
	for _, d := range data {
		floor, err := wrappingPaperNeeded(d.text)
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
