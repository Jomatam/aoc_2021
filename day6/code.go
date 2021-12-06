package day6

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

const fallback = "day6/input.txt"

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	return run(lines)
}

func run(lines []string) (int, int) {
	one := runInner(lines, 80)
	two := runInner(lines, 256)

	fmt.Println(one)
	fmt.Println(two)
	return one, two
}

func runInner(lines []string, iters int) int {
	items := strings.Split(lines[0], ",")

	gen := make([]int, 9)
	for _, item := range items {
		n, _ := strconv.Atoi(item)
		gen[n]++
	}

	for i := 0; i < iters; i++ {
		nextGen := make([]int, 9)
		for i, pop := range gen {
			if i == 0 {
				nextGen[6] = pop
				nextGen[8] = pop
			} else {
				nextGen[i-1] += pop
			}
		}
		gen = nextGen
	}

	result := 0
	for _, pop := range gen {
		result += pop
	}

	return result
}
