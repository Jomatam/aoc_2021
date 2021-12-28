package day6

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"3,4,3,1,2"}
	expected1 := 5934
	expected2 := 26984457539

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
