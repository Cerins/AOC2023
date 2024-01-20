package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	number  int
	winning [100]bool
	actual  [100]bool
}

func parseLine(lines []string) []card {
	// This parsing only works after changing the input txt file
	// I did not bother with adding logic to properly remove additional ""
	return Map(lines, func(line string, _ int) card {
		lineSpl := strings.Split(line, ": ")
		var card card
		card.number, _ = strconv.Atoi(strings.Split(lineSpl[0], " ")[1])
		numbers := strings.Split(lineSpl[1], " | ")
		winning := numbers[0]
		actual := numbers[1]
		for _, w := range strings.Split(winning, " ") {
			if w == "" {
				continue
			}
			parsed, _ := strconv.Atoi(w)
			card.winning[parsed] = true
		}
		for _, a := range strings.Split(actual, " ") {
			if a == "" {
				continue
			}
			parsed, _ := strconv.Atoi(a)
			card.actual[parsed] = true
		}
		return card
	})
}

func solve1(lines []string) {
	cards := parseLine(lines)
	score := 0
	for _, card := range cards {
		found := false
		res := 1
		for i, act := range card.actual {
			if act && card.winning[i] {
				fmt.Println("Found winning at", card.number, i)
				found = true
				res *= 2
			}
		}
		if found {
			res /= 2
			fmt.Println("Total score for ", card.number, "is ", res)
			score += res
		}
	}
	fmt.Println(score)
}

func solve2(lines []string) {
	cards := parseLine(lines)
	queue := make([]int, 0)
	for _, card := range cards {
		queue = append(queue, card.number-1)
	}
	totalWon := 0
	// Maybe do not process every item once but keep a tree structure?
	// To calc the total number of visits
	// Or maybe even better DP
	for len(queue) > 0 {
		totalWon++
		top := queue[0]
		queue = queue[1:]
		// fmt.Println("Reading card", top)
		card := cards[top]
		res := 0
		for i, act := range card.actual {
			if act && card.winning[i] {
				// fmt.Println("Found winning at", card.number, i)
				res++
			}
		}
		for res > 0 {
			queue = append(queue, card.number-1+res)
			res--
		}
	}
	fmt.Println(totalWon)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	solve2(lines)
}
