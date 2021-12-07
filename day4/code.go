package day4

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

const path = "day4/input.txt"

type card = []map[int]bool

func Run() (int, int) {
	lines := helpers.Getlines(path)
	return run(lines)
}

func run(lines []string) (int, int) {
	numbers := getnumbers(lines[0])
	cards := getcards(lines)

	score1 := getScoreForFinishedCard(cards, numbers, 1)
	score2 := getScoreForFinishedCard(cards, numbers, len(cards))

	fmt.Println(score1)
	fmt.Println(score2)
	return score1, score2
}

func getScoreForFinishedCard(cards []card, numbers []int, target int) int {
	finished := make(map[int]bool)
	for _, number := range numbers {
		for i, card := range cards {
			done := false
			for _, row := range card {
				delete(row, number)
				if len(row) == 0 {
					done = true
				}
			}
			if done {
				finished[i] = true
			}
			if len(finished) == target {
				return calculateScore(card, number)
			}
		}
	}
	return -1
}

func calculateScore(card card, number int) int {
	sum := 0
	for _, row := range card {
		for n := range row {
			sum += n
		}
	}
	return sum * number / 2
}

func getnumbers(line string) []int {
	var numbers []int
	items := strings.Split(line, ",")
	for _, item := range items {
		number, _ := strconv.Atoi(item)
		numbers = append(numbers, number)
	}
	return numbers
}

func getcards(lines []string) []card {
	var cards []card

	nofcards := len(lines) / 6

	for i := 0; i < nofcards; i++ {
		cardLines := lines[(i*6 + 2):(i*6 + 7)]
		cards = append(cards, getcard(cardLines))
	}

	return cards
}

func getcard(lines []string) card {
	f := func(c rune) bool { return c == ' ' }

	card := make(card, 10)
	for i := 0; i < 10; i++ {
		card[i] = make(map[int]bool)
	}

	for row, line := range lines {
		items := strings.FieldsFunc(line, f)
		for column, item := range items {
			number, _ := strconv.Atoi(item)
			card[row][number] = true
			card[column+5][number] = true
		}
	}
	return card
}
