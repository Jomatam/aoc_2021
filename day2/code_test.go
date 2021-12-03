package day2

import (
	"testing"
)

func TestResult(t *testing.T) {
	part1, part2 := Run("")

	if part1 != 2215080 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 1864715580 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
