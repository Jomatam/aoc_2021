package day7

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

const path = "day7/input.txt"

func Run() (int, int) {
	lines := helpers.Getlines(path)
	return run(lines)
}

func run(lines []string) (int, int) {
	one := runInner(lines, run1)
	two := runInner(lines, run2)

	fmt.Println(one)
	fmt.Println(two)
	return one, two
}

func runInner(lines []string, f func(int, int) int) int {
	items := strings.Split(lines[0], ",")
	numbers := make([]int, len(items))
	min := int(^uint(0) >> 1)
	max := -min - 1
	for i, item := range items {
		number, _ := strconv.Atoi(item)
		numbers[i] = number
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	minSum := int(^uint(0) >> 1)

	for i := min; i < max; i++ {
		sum := 0
		for _, number := range numbers {
			if i < number {
				sum += f(i, number)
			} else {
				sum += f(number, i)
			}
		}
		if sum < minSum {
			minSum = sum
		}
	}

	return minSum
}

func run1(lower int, upper int) int {
	return upper - lower
}

func run2(lower int, upper int) int {
	return (upper - lower) * (upper - lower + 1) / 2
}
