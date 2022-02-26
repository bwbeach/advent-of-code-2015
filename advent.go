package main

import (
	"errors"
	"fmt"
)

// computeFloor calculates the final floor of the elevator.
//
// An opening parenthesis, (, means he should go up one floor,
// and a closing parenthesis, ), means he should go down one floor.
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
