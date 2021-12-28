package day10

import (
	"aoc_2021/helpers"
	"fmt"
	"sort"
)

const path = "day10/input.txt"

var matchingOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<'}

var matchingCloser = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>'}

var run1Values = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137}

var run2Values = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4}

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
		_, score := step1(line)
		sum += score
	}
	return sum
}

func run2(lines []string) int {
	var values []int
	for _, line := range lines {
		stack, _ := step1(line)
		if len(stack) > 0 {
			values = append(values, step2(stack))
		}
	}

	sort.Ints(values)
	return values[len(values)/2]
}

func step1(line string) (RuneStack, int) {
	stack := NewRuneStack()
	for _, char := range line {
		if opener, ok := matchingOpener[char]; ok {
			newStack, popped := stack.Pop()
			stack = newStack
			if popped != opener {
				return NewRuneStack(), run1Values[char]
			}
		} else {
			stack = stack.Push(char)
		}
	}
	return stack, 0
}

func step2(stack RuneStack) int {
	result := 0
	for len(stack) > 0 {
		newStack, opener := stack.Pop()
		result *= 5
		result += run2Values[matchingCloser[opener]]
		stack = newStack
	}
	return result
}
