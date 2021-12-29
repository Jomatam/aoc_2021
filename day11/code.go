package day11

import (
	"aoc_2021/helpers"
	"fmt"
)

const path = "day11/input.txt"

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
	grid := helpers.NewIntGrid(lines)
	flashes := 0
	for i := 0; i < 100; i++ {
		incrementAll(&grid)
		flashAll(&grid)
		flashes += resetFlashes(&grid)
	}
	return flashes
}

func run2(lines []string) int {
	grid := helpers.NewIntGrid(lines)
	iterations := 0
	for {
		result := resetFlashes(&grid)
		if result == 100 {
			return iterations
		}
		incrementAll(&grid)
		flashAll(&grid)
		iterations++
	}
}

func incrementAll(grid *helpers.IntGrid) {
	for y, xs := range *grid {
		for x := range xs {
			increment(grid, x, y)
		}
	}
}

func increment(grid *helpers.IntGrid, x, y int) {
	if original, ok := grid.Get(x, y); ok {
		grid.Set(x, y, original+1)
	}
}

func reset(grid *helpers.IntGrid, x, y int) {
	if _, ok := grid.Get(x, y); ok {
		grid.Set(x, y, 0)
	}
}

func flashAll(grid *helpers.IntGrid) {
	flashGrid := helpers.NewBoolGrid(len((*grid)[0]), len(*grid), false)
	anyFlash := true
	for anyFlash {
		anyFlash = false
		for y, xs := range *grid {
			for x, value := range xs {
				alreadyFlashed, _ := flashGrid.Get(x, y)
				if !alreadyFlashed && value > 9 {
					anyFlash = true
					increment(grid, x-1, y-1)
					increment(grid, x-1, y)
					increment(grid, x-1, y+1)
					increment(grid, x, y-1)
					increment(grid, x, y+1)
					increment(grid, x+1, y-1)
					increment(grid, x+1, y)
					increment(grid, x+1, y+1)
					flashGrid.Set(x, y, true)
				}
			}
		}
	}
}

func resetFlashes(grid *helpers.IntGrid) int {
	result := 0
	for y, xs := range *grid {
		for x, value := range xs {
			if value > 9 {
				result++
				reset(grid, x, y)
			}
		}
	}
	return result
}
