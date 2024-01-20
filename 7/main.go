package main

import (
	. "AOC2023/utils"

	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var importance = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

type game struct {
	cards string
	score int
}

func parseLines(lines []string) []game {
	return Map(lines, func(line string, i int) game {
		var game game
		strs := strings.Split(line, " ")
		game.cards = strs[0]
		game.score, _ = strconv.Atoi(strs[1])
		return game
	})

}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func cardScore(cards string) int {
	count := make([]int, len(importance))
	for _, c := range cards {
		index := IndexOf(c, importance)
		count[index]++
	}
	jIndex := IndexOf('J', importance)
	jCount := count[jIndex]
	count[jIndex] = 0
	sbest := -1
	best := -1
	for _, k := range count {
		if k > best {
			if best > sbest {
				sbest = best
			}
			best = k
		} else if k > sbest {
			sbest = k
		}
	}
	best += jCount
	// Assume that jCounts are all used up
	jCount = 0
	// but since best can overflow
	if best > len(cards) {
		// then jCount is the remainder
		jCount = best - len(cards)
		// and best is the length
		best = len(cards)
	}
	sbest += jCount
	// fmt.Println(cards, best, sbest)
	return best*10 + sbest
}

type sortGames []game

func (s sortGames) Less(i, j int) bool {
	ci := cardScore(s[i].cards)
	cj := cardScore(s[j].cards)
	if ci == cj {
		for k := 0; k < len(s[i].cards); k++ {
			ii := IndexOf(rune(s[i].cards[k]), importance)
			ij := IndexOf(rune(s[j].cards[k]), importance)
			if ii != ij {
				return ii > ij
			}
		}
	}
	return ci < cj
}

func (s sortGames) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortGames) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func solve1(lines []string) int {
	res := 0
	games := parseLines(lines)
	sort.Sort(sortGames(games))
	// fmt.Println(games)
	for i, game := range games {
		res += (i + 1) * game.score
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}
	// fmt.Println(lines)
	input := solve1(lines)
	fmt.Println(input)
}
