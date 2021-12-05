package day5

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

const fallback = "day5/input.txt"

type key struct {
	X, Y int
}

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	return run(lines)
}

func run(lines []string) (int, int) {
	part1 := runInner(lines, false)
	part2 := runInner(lines, true)

	fmt.Println(part1)
	fmt.Println(part2)

	return part1, part2
}

func runInner(lines []string, includeDiagonals bool) int {
	points := make(map[key]int)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		sparts := strings.Split(parts[0], ",")
		fparts := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(sparts[0])
		y1, _ := strconv.Atoi(sparts[1])
		x2, _ := strconv.Atoi(fparts[0])
		y2, _ := strconv.Atoi(fparts[1])

		if x1 == x2 || y1 == y2 {
			for _, x := range rangeInclusive(x1, x2) {
				for _, y := range rangeInclusive(y1, y2) {
					points[key{x, y}]++
				}
			}
		} else if includeDiagonals {
			xs := rangeInclusive(x1, x2)
			ys := rangeInclusive(y1, y2)
			for i, x := range xs {
				points[key{x, ys[i]}]++
			}
		}
	}

	result := 0
	for _, value := range points {
		if value > 1 {
			result++
		}
	}

	return result
}

func rangeInclusive(a int, b int) []int {
	var result []int
	direction := 1
	if b < a {
		direction = -1
	}

	for i := a; i != b; i += direction {
		result = append(result, i)
	}

	result = append(result, b)

	return result
}
