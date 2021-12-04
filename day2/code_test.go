package day2

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	part1, part2 := run(input)

	if part1 != 150 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 900 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
