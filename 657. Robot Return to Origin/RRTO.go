package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, item := range moves {
		switch item {
		case 'U':
			y++
			break
		case 'D':
			y--
			break
		case 'L':
			x--
			break
		case 'R':
			x++
			break
		}

	}

	return x == 0 && y == 0
}

//Best ans
func Bestans(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "L") == strings.Count(moves, "R")
}
func main() {
	fmt.Print(judgeCircle("UD"))
}
