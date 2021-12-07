package infi

import (
	"aoc_2021/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const path = "infi/input.txt"

func Run() {
	lines := helpers.Getlines(path)
	things := getthings(lines)
	amounts := getamounts(things)
	missing, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
	keys := getamountkeys(amounts)

	var empty []int
	recurse(keys, missing, empty, 0)
	converttosequence(amounts, keys)
}

func converttosequence(amounts map[int]string, keys []int) {
	sequence := make([]string, len(keys))
	for i, key := range keys {
		sequence[i] = string(amounts[key][0])
	}
	sort.Strings(sequence)
	fmt.Println(strings.Join(sequence, ""))
}

func getamountkeys(amounts map[int]string) []int {
	var keys []int
	for amount := range amounts {
		keys = append(keys, amount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func getthings(lines []string) map[string]map[string]int {
	things := make(map[string]map[string]int)
	for _, line := range lines[1:] {
		split := strings.Split(line, ": ")
		key := split[0]
		parts := strings.Split(split[1], ", ")
		things[key] = make(map[string]int)
		for _, part := range parts {
			p := strings.Split(part, " ")
			k := p[1]
			v, _ := strconv.Atoi(p[0])
			things[key][k] = v
		}
	}
	return things
}

func getamounts(things map[string]map[string]int) map[int]string {
	amounts := make(map[int]string)
	for thing := range things {
		amount := getamount(things, thing)
		amounts[amount] = thing
	}

	for amount, name := range amounts {
		for _, parts := range things {
			for part := range parts {
				if name == part {
					delete(amounts, amount)
				}
			}
		}
	}
	return amounts
}

func getamount(things map[string]map[string]int, thing string) int {
	parts := things[thing]
	total := 0
	for part, amount := range parts {
		if _, ok := things[part]; ok {
			total += amount * getamount(things, part)
		} else {
			total += amount
		}
	}
	return total
}

func recurse(keys []int, missing int, partial []int, sum int) []int {
	if sum == missing {
		return partial
	}

	if len(partial) > 20 {
		return nil
	}

	for _, amount := range keys {
		result := recurse(keys, missing, append(partial, amount), sum+amount)
		if result != nil {
			return result
		}
	}

	return nil
}
