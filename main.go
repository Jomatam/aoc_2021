package main

import (
	"aoc_2021/day1"
	"aoc_2021/day10"
	"aoc_2021/day2"
	"aoc_2021/day3"
	"aoc_2021/day4"
	"aoc_2021/day5"
	"aoc_2021/day6"
	"aoc_2021/day7"
	"aoc_2021/day8"
	"aoc_2021/day9"
	"aoc_2021/infi"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	case "3":
		day3.Run()
	case "4":
		day4.Run()
	case "5":
		day5.Run()
	case "6":
		day6.Run()
	case "7":
		day7.Run()
	case "8":
		day8.Run()
	case "9":
		day9.Run()
	case "10":
		day10.Run()
	case "i":
		infi.Run()
	}
}
