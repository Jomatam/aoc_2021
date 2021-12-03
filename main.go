package main

import (
	"aoc_2021/day1"
	"aoc_2021/day2"
	"aoc_2021/day3"
	"os"
	"strconv"
)

func main() {
	run, _ := strconv.Atoi(os.Args[1])

	switch run {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 3:
		day3.Run()
	}
}
