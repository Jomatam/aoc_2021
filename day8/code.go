package day8

import (
	"aoc_2021/helpers"
	"fmt"
	"math"
	"strings"
)

const path = "day8/input.txt"

var values = []string{"a", "b", "c", "d", "e", "f", "g"}

type Digit struct {
	left_upper  string
	left_lower  string
	right_lower string
}

func newDigit() Digit {
	return Digit{
		left_upper:  "",
		left_lower:  "",
		right_lower: ""}
}

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
	sum := 0
	for _, line := range lines {
		relevant := strings.Split(line, " | ")[1]
		parts := strings.Split(relevant, " ")
		for _, part := range parts {
			if len(part) == 2 ||
				len(part) == 3 ||
				len(part) == 4 ||
				len(part) == 7 {
				sum++
			}
		}
	}
	return sum
}

func run2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += process(line)
	}

	return sum
}

func process(line string) int {
	split := strings.Split(line, " | ")
	outputs := strings.Split(split[1], " ")

	digit := newDigit()
	for _, r := range values {
		count := strings.Count(split[0], r)
		switch count {
		case 4:
			digit.left_lower = r
		case 6:
			digit.left_upper = r
		case 9:
			digit.right_lower = r
		}
	}
	one := ""
	for _, r := range strings.Split(split[0], " ") {
		if len(r) == 2 {
			one = r
		}
	}

	result := 0
	for i, output := range outputs {
		factor := int(math.Pow10(3 - i))
		if len(output) == 2 {
			result += 1 * factor
		} else if len(output) == 3 {
			result += 7 * factor
		} else if len(output) == 4 {
			result += 4 * factor
		} else if len(output) == 7 {
			result += 8 * factor
		} else if !strings.Contains(output, digit.right_lower) {
			result += 2 * factor
		} else if len(output) == 5 && strings.Contains(output, digit.left_upper) {
			result += 5 * factor
		} else if len(output) == 5 {
			result += 3 * factor
		} else if !strings.Contains(output, digit.left_lower) {
			result += 9 * factor
		} else if !containsone(output, one) {
			result += 6 * factor
		}
	}
	return result
}

func containsone(digit string, one string) bool {
	for _, c := range one {
		if !strings.Contains(digit, string(c)) {
			return false
		}
	}
	return true
}
