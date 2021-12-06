package day6

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"3,4,3,1,2"}

	part1, part2 := run(input)

	if part1 != 5934 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 26984457539 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
