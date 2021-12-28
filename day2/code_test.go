package day2

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expected1 := 150
	expected2 := 900

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
