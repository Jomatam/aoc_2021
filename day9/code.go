package day9

import (
	"aoc_2021/helpers"
	"fmt"
	"sort"
	"strconv"
)

const path = "day9/input.txt"

func Run() (int, int) {
	lines := helpers.Getlines(path)
	return run(lines)
}

func run(lines []string) (int, int) {
	one := run1(lines)
	two := run2(lines)

	fmt.Println(one)
	fmt.Println(two)
	return one, two
}

func run1(lines []string) int {
	grid := getGrid(lines)
	sum := 0
	for y, xs := range grid {
		for x, value := range xs {
			if isLowPoint(&grid, x, y) {
				sum += value + 1
			}
		}
	}
	return sum
}

func run2(lines []string) int {
	grid := getGrid(lines)
	var sizes []int
	x, y := 0, 0
	for x+y >= 0 {
		sizes = append(sizes, fill(&grid, x, y))
		x, y = getNextBasin(&grid)
	}

	sort.Ints(sizes)
	result := 1
	for _, value := range sizes[len(sizes)-3:] {
		result *= value
	}
	return result
}

func fill(grid *[][]int, x, y int) int {
	if getHeight(grid, x, y) > 8 {
		return 0
	} else {
		setHeight(grid, x, y, 9)
		return 1 +
			fill(grid, x+1, y) +
			fill(grid, x-1, y) +
			fill(grid, x, y+1) +
			fill(grid, x, y-1)
	}
}

func getNextBasin(grid *[][]int) (int, int) {
	for y, xs := range *grid {
		for x, value := range xs {
			if value < 9 {
				return x, y
			}
		}
	}
	return -1, -1
}

func getGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			grid[y][x], _ = strconv.Atoi(string(char))
		}
	}
	return grid
}

func isLowPoint(grid *[][]int, x, y int) bool {
	this := getHeight(grid, x, y)
	return this < getHeight(grid, x+1, y) &&
		this < getHeight(grid, x-1, y) &&
		this < getHeight(grid, x, y+1) &&
		this < getHeight(grid, x, y-1)
}

func setHeight(grid *[][]int, x, y, height int) {
	(*grid)[y][x] = height
}

func getHeight(grid *[][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len((*grid)[0]) || y >= len(*grid) {
		return 10
	}
	return (*grid)[y][x]
}
