package day1

import (
	"testing"
)

func TestResult(t *testing.T) {
	part1, part2 := Run("")

	if part1 != 1233 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 1275 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
