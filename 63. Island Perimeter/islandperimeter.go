package main

import (
	"fmt"
)

// You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

// Grid cells are connected horizontally/vertically (not diagonally).
// The grid is completely surrounded by water, and there is exactly one island
// (i.e., one or more connected land cells).

// The island doesn't have "lakes" (water inside that isn't connected to the water around the island).
// One cell is a square with side length 1.
// The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

func islandPerimeter(grid [][]int) int {
	res := 0

	for idx := 0; idx < len(grid); idx++ {

		for idx2 := 0; idx2 < len(grid[0]); idx2++ {
			fmt.Println(grid[idx][idx2])
			if grid[idx][idx2] == 1 {

				if grid[idx][idx2] == 1 {
					//up
					if idx == 0 {
						res++
					} else {
						if grid[idx-1][idx2] == 0 {
							res++
						}
					}
					//down
					if idx == len(grid)-1 {
						res++
					} else {
						if grid[idx+1][idx2] == 0 {
							res++
						}
					}

					//right
					if idx2 == len(grid[0])-1 {
						res++
					} else {
						if grid[idx][idx2+1] == 0 {
							res++
						}
					}
					//left
					if idx2 == 0 {
						res++
					} else {
						if grid[idx][idx2-1] == 0 {
							res++
						}
					}

				}
				//fmt.Printf("%d %d\n", idx, idx2)
			}
		}
	}
	fmt.Printf("res : %d\n", res)
	return res
}

func main() {
	islandPerimeter([][]int{
		{0, 1},
	})

}
