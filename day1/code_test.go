package day1

import (
	"testing"
)

func TestResult(t *testing.T) {
	var input = []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	part1, part2 := run(input)

	if part1 != 7 {
		t.Fatalf("Run with default input failed at part 1")
	}

	if part2 != 5 {
		t.Fatalf("Run with default input failed at part 2")
	}
}
