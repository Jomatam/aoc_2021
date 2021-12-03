package main

import (
	"aoc_2021/day1"
	"aoc_2021/day2"
	"aoc_2021/day3"
	"aoc_2021/infi"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	case "3":
		day3.Run()
	case "i":
		infi.Run()
	}
}
