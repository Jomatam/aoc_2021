package day3

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
)

var bits int

const fallback = "day3/input.txt"

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	return run(lines)
}

func run(lines []string) (int, int) {
	bits = len(lines[0])

	ints := make([]int, len(lines))
	for i, line := range lines {
		x, _ := strconv.ParseInt(line, 2, 64)
		ints[i] = int(x)
	}

	g := mostLeastCommon(ints, true)
	e := mostLeastCommon(ints, false)
	fmt.Println(g * e)

	o := getMostSimilar(ints, true)
	c := getMostSimilar(ints, false)
	fmt.Println(o * c)

	return g * e, o * c
}

func getMostSimilar(ints []int, most bool) int {
	for i := 0; i < bits; i++ {
		t := mostLeastCommon(ints, most)
		ints = filter(ints, t, i)
		if len(ints) == 1 {
			break
		}
	}
	return ints[0]
}

func mostLeastCommon(ints []int, most bool) int {
	counts := make([]int, bits)
	for bit := 0; bit < bits; bit++ {
		mask := 1 << (bits - bit - 1)
		for _, i := range ints {
			if i&mask > 0 {
				counts[bit]++
			} else {
				counts[bit]--
			}
		}
	}

	common := 0
	for i, value := range counts {
		mask := 1 << (bits - i - 1)
		if (value >= 0) == most {
			common |= mask
		}
	}
	return common
}

func filter(values []int, target int, index int) []int {
	var filtered []int
	mask := 1 << (bits - index - 1)
	for _, value := range values {
		if (value&mask == 0) == (target&mask == 0) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}
