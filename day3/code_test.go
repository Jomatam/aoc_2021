package day3

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"00100", "11110", "10110", "10111", "10101",
		"01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected1 := 198
	expected2 := 230

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
