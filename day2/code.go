package day2

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	lines := helpers.Getlines("day2/input.txt")
	part_1(lines)
	part_2(lines)
}

func part_1(lines []string) {
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

	fmt.Println(x * y)
}

func part_2(lines []string) {
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

	fmt.Println(x * y)
}
