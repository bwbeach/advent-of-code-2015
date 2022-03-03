package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
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

// parsePackage returns the dimensions, smallest first, of
// a package, given a package spec string like "4x2x12"
func parsePackage(spec string) ([]int, error) {
	words := strings.Split(spec, "x")
	if len(words) != 3 {
		return nil, errors.New("bad package spec")
	}

	dims := make([]int, 3)
	for i := 0; i < 3; i++ {
		dim, err := strconv.Atoi(words[i])
		if err != nil {
			return nil, err
		}
		dims[i] = dim
	}

	sort.Ints(dims)

	return dims, nil
}

// wrappingPaperNeeded returns the number of square feet needed
// for a package, given the dimensions.
func wrappingPaperNeeded(dims []int) int {
	return 3*dims[0]*dims[1] + 2*dims[0]*dims[2] + 2*dims[1]*dims[2]
}

// ribbonNeeded returns the number of feet of ribbon needed for
// a package, given the dimensions
func ribbonNeeded(dims []int) int {
	return 2*dims[0] + 2*dims[1] + dims[0]*dims[1]*dims[2]
}

func day02() {
	file, err := os.Open("day02-input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	totalPaper := 0
	totalRibbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dims, err := parsePackage(scanner.Text())
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err.Error())
			os.Exit(1)
		}

		totalPaper += wrappingPaperNeeded(dims)
		totalRibbon += ribbonNeeded(dims)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Day 2")
	fmt.Println("Total paper needed: ", totalPaper)
	fmt.Println("Total ribbon needed: ", totalRibbon)
}

type point struct {
	x int
	y int
}

func nextHouse(start point, move rune) (point, error) {
	switch move {
	case '<':
		return point{x: start.x - 1, y: start.y}, nil
	case '>':
		return point{x: start.x + 1, y: start.y}, nil
	case '^':
		return point{x: start.x, y: start.y + 1}, nil
	case 'v':
		return point{x: start.x, y: start.y - 1}, nil
	default:
		return point{}, errors.New("bad move")
	}
}

// housesForMoves returns the set of houses visit for
// one string of moves, started at (0, 0)
func housesForMoves(moves string) (map[point]bool, error) {
	houses := make(map[point]bool, 0)
	loc := point{x: 0, y: 0}
	var err error
	houses[loc] = true
	for _, c := range string(moves) {
		loc, err = nextHouse(loc, c)
		if err != nil {
			return nil, fmt.Errorf("Bad input char %v", loc)
		}
		houses[loc] = true
	}
	return houses, nil
}

func day03() {
	fmt.Println("Day 3")
	text, err := ioutil.ReadFile("day03-input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err.Error())
		os.Exit(1)
	}

	houses, err := housesForMoves(string(text))
	if err != nil {
		fmt.Println("Error moving: ", err)
		os.Exit(1)
	}

	houseCount := len(houses)
	fmt.Println("Number of houses: ", houseCount)
}

func main() {

	fmt.Println("Day 1")
	text, err := ioutil.ReadFile("day01-input.txt")
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

	day02()
	day03()
}
