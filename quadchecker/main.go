package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Quad struct {
	name  string
	check func(lines []string, width, height int) bool
}

var quads = []Quad{
	{"quadA", checkQuadA},
	{"quadB", checkQuadB},
	{"quadC", checkQuadC},
	{"quadD", checkQuadD},
	{"quadE", checkQuadE},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	width := len(lines[0])
	height := len(lines)

	// Check all lines have same width
	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	var matches []string
	for _, quad := range quads {
		if quad.check(lines, width, height) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad.name, width, height))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		sort.Strings(matches)
		fmt.Println(strings.Join(matches, " || "))
	}
}

// Quad A: o--o
//
// | |
// o--o
func checkQuadA(lines []string, width, height int) bool {
	if width < 1 || height < 1 {
		return false
	}
	for y, line := range lines {
		for x, char := range line {
			switch {
			case (y == 0 || y == height-1) && (x == 0 || x == width-1):
				if char != 'o' {
					return false
				}
			case y == 0 || y == height-1:
				if char != '-' {
					return false
				}
			case x == 0 || x == width-1:
				if char != '|' {
					return false
				}
			default:
				if char != ' ' {
					return false
				}
			}
		}
	}
	return true
}

// Quad B: /***\
// - *
// \***/
func checkQuadB(lines []string, width, height int) bool {
	if width < 1 || height < 1 {
		return false
	}
	for y, line := range lines {
		for x, char := range line {
			switch {
			case y == 0 && x == 0:
				if char != '/' {
					return false
				}
			case y == 0 && x == width-1:
				if char != '\\' {
					return false
				}
			case y == height-1 && x == 0:
				if char != '\\' {
					return false
				}
			case y == height-1 && x == width-1:
				if char != '/' {
					return false
				}
			case y == 0 || y == height-1:
				if char != '*' {
					return false
				}
			case x == 0 || x == width-1:
				if char != '*' {
					return false
				}
			default:
				if char != ' ' {
					return false
				}
			}
		}
	}
	return true
}

// Quad C: ABBA
//
// B B
// C C
func checkQuadC(lines []string, width, height int) bool {
	if width < 1 || height < 1 {
		return false
	}
	for y, line := range lines {
		for x, char := range line {
			switch {
			case y == 0 && (x == 0 || x == width-1):
				if char != 'A' {
					return false
				}
			case y == height-1 && (x == 0 || x == width-1):
				if char != 'C' {
					return false
				}
			case y == 0 || y == height-1:
				if char != 'B' {
					return false
				}
			case x == 0 || x == width-1:
				if char != 'B' {
					return false
				}
			default:
				if char != ' ' {
					return false
				}
			}
		}
	}
	return true
}

// Quad D: ABBC
//
// B B
// ABBC
func checkQuadD(lines []string, width, height int) bool {
	if width < 1 || height < 1 {
		return false
	}
	for y, line := range lines {
		for x, char := range line {
			switch {
			case (y == 0 || y == height-1) && x == 0:
				if char != 'A' {
					return false
				}
			case (y == 0 || y == height-1) && x == width-1:
				if char != 'C' {
					return false
				}
			case y == 0 || y == height-1:
				if char != 'B' {
					return false
				}
			case x == 0 || x == width-1:
				if char != 'B' {
					return false
				}
			default:
				if char != ' ' {
					return false
				}
			}
		}
	}
	return true
}

// Quad E: ABBC
//
// B B
// CBB
func checkQuadE(lines []string, width, height int) bool {
	if width < 1 || height < 1 {
		return false
	}
	for y, line := range lines {
		for x, char := range line {
			switch {
			case y == 0 && x == 0:
				if char != 'A' {
					return false
				}
			case y == 0 && x == width-1:
				if char != 'C' {
					return false
				}
			case y == height-1 && x == 0:
				if char != 'C' {
					return false
				}
			case y == height-1 && x == width-1:
				if char != 'A' {
					return false
				}
			case y == 0 || y == height-1:
				if char != 'B' {
					return false
				}
			case x == 0 || x == width-1:
				if char != 'B' {
					return false
				}
			default:
				if char != ' ' {
					return false
				}
			}
		}
	}
	return true
}
