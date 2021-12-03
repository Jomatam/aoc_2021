package day1

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
)

func Run() {
	lines := helpers.Getlines("day1/input.txt")
	var values []int
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}

	fmt.Println(getincreases(values, 1))
	fmt.Println(getincreases(values, 3))
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
