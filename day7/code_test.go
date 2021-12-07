package day7

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"16,1,2,0,4,2,7,1,2,14"}

	part1, part2 := run(input)

	if part1 != 37 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 168 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
