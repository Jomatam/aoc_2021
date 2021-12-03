package day3

import (
	"testing"
)

func TestResult(t *testing.T) {
	part1, part2 := Run("")

	if part1 != 3813416 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 2990784 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
