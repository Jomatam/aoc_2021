package day9

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678"}
	expected1 := 15
	expected2 := 1134

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
