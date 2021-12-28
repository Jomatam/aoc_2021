package day1

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	expected1 := 7
	expected2 := 5

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
