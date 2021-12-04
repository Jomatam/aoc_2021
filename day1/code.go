package day1

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
)

const fallback = "day1/input.txt"

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	return run(lines)
}

func run(lines []string) (int, int) {
	values := make([]int, len(lines))
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		values[i] = value
	}

	one := getNumberOfIncreases(values, 1)
	three := getNumberOfIncreases(values, 3)

	fmt.Println(one)
	fmt.Println(three)

	return one, three
}

func getNumberOfIncreases(values []int, window int) int {
	result := 0
	for i, value := range values[window:] {
		if value > values[i] {
			result++
		}
	}
	return result
}
