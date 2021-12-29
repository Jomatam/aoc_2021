package day12

import (
	"aoc_2021/helpers"
	"fmt"
	"strings"
	"unicode"
)

const path = "day12/input.txt"

func Run() (int, int) {
	lines := helpers.Getlines(path)
	return run(lines)
}

func run(lines []string) (int, int) {
	one := run1(lines)
	two := run2(lines)

	fmt.Println(one)
	fmt.Println(two)
	return one, two
}

func run1(lines []string) int {
	return runInner(lines, false)
}

func run2(lines []string) int {
	return runInner(lines, true)
}

// using a simple and specific implementation of bfs
func runInner(lines []string, allowDouble bool) int {
	connections := getConnections(lines)
	var queue helpers.StringQueue
	queue = queue.Enqueue("start")
	routes := make(map[string]bool)
	result := 0

	for {
		newQueue, partial, ok := queue.Dequeue()
		if !ok {
			return result
		}
		queue = newQueue
		if _, found := routes[partial]; found {
			continue
		}
		routes[partial] = true
		if strings.HasSuffix(partial, "end") {
			result++
			continue
		}
		caves := strings.Split(partial, ",")
		lastCave := caves[len(caves)-1]
		for _, nextCave := range connections[lastCave] {
			if newPartial, ok := nextPartial(partial, nextCave, caves, allowDouble); ok {
				queue = queue.Enqueue(newPartial)
			}
		}
	}
}

func nextPartial(partial, nextCave string, caves []string, allowDouble bool) (string, bool) {
	newPartial := partial + "," + nextCave

	if unicode.IsUpper(rune(nextCave[0])) {
		return newPartial, true
	}

	if nextCave == "start" {
		return "", false
	}

	for _, cave := range caves {
		if cave == nextCave {
			// Add an _ to determine whether a route already has a double visit to a small cave.
			// Otherwise we would need to have a nexted loop over the caves, or complicate the queue.
			if !allowDouble || strings.HasPrefix(partial, "_") {
				return "", false
			} else {
				return "_" + newPartial, true
			}
		}
	}
	return newPartial, true
}

func getConnections(lines []string) map[string][]string {
	result := make(map[string][]string)
	for _, line := range lines {
		caves := strings.Split(line, "-")
		result[caves[0]] = append(result[caves[0]], caves[1])
		result[caves[1]] = append(result[caves[1]], caves[0])
	}
	return result
}
