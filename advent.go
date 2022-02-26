package main

import (
	"errors"
	"fmt"
)

func computeFloor(text string) (int, error) {
	floor := 0
	for _, c := range text {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		} else {
			return 0, errors.New("bad character")
		}
	}
	return floor, nil
}

func main() {
	fmt.Println("hello")
}
