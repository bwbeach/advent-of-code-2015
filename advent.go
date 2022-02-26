package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// nextFloor returns the floor after processing one instruction
func nextFloor(prevFloor int, instruction rune) (int, error) {
	if instruction == '(' {
		return prevFloor + 1, nil
	} else if instruction == ')' {
		return prevFloor - 1, nil
	} else {
		return 0, errors.New("bad character")
	}
}

// computeFloor calculates the final floor of the elevator.
//
// An opening parenthesis, (, means he should go up one floor,
// and a closing parenthesis, ), means he should go down one floor.
func computeFloor(text string) (int, error) {
	floor := 0
	var err error
	for _, c := range text {
		floor, err = nextFloor(floor, c)
		if err != nil {
			return 0, err
		}
	}
	return floor, nil
}

// basementPosition returns the index (one-based) of the first instruction
// that puts the elevator in the basement.
func basementPosition(text string) (int, error) {
	floor := 0
	var err error
	for i, c := range text {
		floor, err = nextFloor(floor, c)
		if err != nil {
			return 0, err
		}
		if floor == -1 {
			return i + 1, nil
		}
	}
	return 0, errors.New("never entered basement")
}

func main() {
	text, err := ioutil.ReadFile("day01_input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		os.Exit(1)
	}

	floor, err := computeFloor(string(text))
	if err != nil {
		fmt.Printf("Error computing floor: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("floor = %d\n", floor)

	basementPos, err := basementPosition(string(text))
	if err != nil {
		fmt.Printf("Error finding basement pos: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("basementPos = %d\n", basementPos)

}
