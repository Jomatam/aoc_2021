package helpers

import "strconv"

type IntGrid [][]int

func NewIntGrid(lines []string) IntGrid {
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			grid[y][x], _ = strconv.Atoi(string(char))
		}
	}
	return grid
}

func (grid *IntGrid) Set(x, y, value int) {
	(*grid)[y][x] = value
}

func (grid *IntGrid) Get(x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= len((*grid)[0]) || y >= len(*grid) {
		return 0, false
	}
	return (*grid)[y][x], true
}
