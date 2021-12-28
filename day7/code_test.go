package day7

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"16,1,2,0,4,2,7,1,2,14"}
	expected1 := 37
	expected2 := 168

	part1, part2 := run(input)

	if part1 != expected1 {
		t.Fatalf("Run with default input failed at part 1. Expected %d, returned %d", expected1, part2)
	}

	if part2 != expected2 {
		t.Fatalf("Run with default input failed at part 2. Expected %d, returned %d", expected2, part2)
	}
}
