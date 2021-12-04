package day3

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"00100", "11110", "10110", "10111", "10101",
		"01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	part1, part2 := run(input)

	if part1 != 198 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 230 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
