package main

import (
	. "AOC2023/utils"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	firstRow []int
}

func parseLines(lines []string) []game {
	return Map(lines, func(line string, i int) game {
		var game game
		strs := strings.Split(line, " ")
		game.firstRow = make([]int, len(strs))
		for i, str := range strs {
			game.firstRow[i], _ = strconv.Atoi(str)
		}
		return game
	})
}

func ncr(n, r int) int {
	if r > n {
		return 0
	}
	if r == 0 || r == n {
		return 1
	}
	res := 1
	for i := 1; i <= r; i++ {
		res *= n - i + 1
		res /= i
	}
	return res
}

func sign(i int) int {
	if i%2 == 0 {
		return 1
	}
	return -1
}

// Todays part one can be solve by ncr
func solve1(lines []string) int {
	res := 0
	games := parseLines(lines)
	for _, game := range games {
		n := len(game.firstRow)
		var sum, i int
		for ; i < n; i++ {
			sum += ncr(n, i) * sign(i) * game.firstRow[i]
		}
		res += (-sum) * sign(i)
	}
	fmt.Println(games)
	return res
}
func solve2(lines []string) int {
	res := 0
	games := parseLines(lines)
	for _, game := range games {
		n := len(game.firstRow)
		var sum, i int
		for i = 1; i <= n; i++ {
			// fmt.Println(ncr(n, i), sign(i), game.firstRow[i], ncr(n, i)*sign(i)*game.firstRow[i])
			sum += ncr(n, i) * sign(i) * game.firstRow[i-1]
		}
		fmt.Println(game.firstRow, sum*sign(i))
		res += -sum
	}
	// fmt.Println(games)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}
	res := solve2(lines)
	fmt.Println(res)
}
