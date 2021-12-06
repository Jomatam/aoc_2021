package main

import (
	"aoc_2021/day1"
	"aoc_2021/day2"
	"aoc_2021/day3"
	"aoc_2021/day4"
	"aoc_2021/day5"
	"aoc_2021/day6"
	"aoc_2021/infi"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	day := os.Args[1]

	var path string
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	switch day {
	case "1":
		day1.Run(path)
	case "2":
		day2.Run(path)
	case "3":
		day3.Run(path)
	case "4":
		day4.Run(path)
	case "5":
		day5.Run(path)
	case "6":
		day6.Run(path)
	case "i":
		infi.Run()
	}
}
