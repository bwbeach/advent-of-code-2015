package main

import (
	"reflect"
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

func Test_parsePackage(t *testing.T) {
	data := []struct {
		text            string
		expected        []int
		expectedMessage string
	}{
		{"2x3x4", []int{2, 3, 4}, ""},
		{"3x4x2", []int{2, 3, 4}, ""},
		{"3x1", nil, "bad package spec"},
		{"3xAx4", nil, "strconv.Atoi: parsing \"A\": invalid syntax"},
	}
	for _, d := range data {
		dims, err := parsePackage(d.text)
		if !reflect.DeepEqual(dims, d.expected) {
			t.Errorf("for '%s' expected %d but got %d", d.text, d.expected, dims)
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

func Test_wrappingPaperNeeded(t *testing.T) {
	data := []struct {
		dims     []int
		expected int
	}{
		{[]int{2, 3, 4}, 58},
	}
	for _, d := range data {
		area := wrappingPaperNeeded(d.dims)
		if area != d.expected {
			t.Errorf("expected %d but got %d", d.expected, area)
		}
	}
}

func Test_ribbonNeeded(t *testing.T) {
	data := []struct {
		dims     []int
		expected int
	}{
		{[]int{2, 3, 4}, 34},
		{[]int{1, 1, 10}, 14},
	}
	for _, d := range data {
		feet := ribbonNeeded(d.dims)
		if feet != d.expected {
			t.Errorf("expected %d but got %d", d.expected, feet)
		}
	}
}
