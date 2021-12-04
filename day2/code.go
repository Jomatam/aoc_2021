package day2

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

const fallback = "day2/input.txt"

func Run(path string) (int, int) {
	lines := helpers.Getlines(path, fallback)
	return run(lines)
}

func run(lines []string) (int, int) {
	part1 := part_1(lines)
	part2 := part_2(lines)

	fmt.Println(part1)
	fmt.Println(part2)

	return part1, part2
}

func part_1(lines []string) int {
	x := 0
	y := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			x += amount
		case "down":
			y += amount

		case "up":
			y -= amount
		}
	}

	return x * y
}

func part_2(lines []string) int {
	x := 0
	y := 0
	aim := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			x += amount
			y += amount * aim
		case "down":
			aim += amount

		case "up":
			aim -= amount
		}
	}

	return x * y
}
