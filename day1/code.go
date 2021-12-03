package day1

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
)

const fallback = "day1/input.txt"

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	var values []int
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}

	one := getincreases(values, 1)
	three := getincreases(values, 3)

	fmt.Println(one)
	fmt.Println(three)

	return one, three
}

func getincreases(values []int, window int) int {
	result := 0
	for i, value := range values[window:] {
		if value > values[i] {
			result++
		}
	}
	return result
}
