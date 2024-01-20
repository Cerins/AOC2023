package main

import (
	. "AOC2023/utils"

	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	time int
	beat int
}

func parseLines(lines []string) []game {
	// Today i will try regex
	r := regexp.MustCompile(`\d+`)
	times := r.FindAllString(lines[0], -1)
	beats := r.FindAllString(lines[1], -1)
	return Map(times, func(time string, i int) game {
		tp, _ := strconv.Atoi(time)
		bp, _ := strconv.Atoi(beats[i])
		return game{
			time: tp,
			beat: bp,
		}
	})

}

func solve1(lines []string) int {
	input := parseLines(lines)
	res := 1
	for _, game := range input {
		time := float64(game.time)
		beat := float64(game.beat)
		z := math.Sqrt(time*time - 4*beat)
		x1 := -(-time + z) / 2
		x2 := -(-time - z) / 2
		a1 := int(math.Floor(x1) + 1)
		a2 := int(math.Ceil(x2) - 1)
		beats := (a2 - a1 + 1)
		// Maybe is possible to not use a1 and a2?
		fmt.Println(-time, z)
		if beats < 0 {
			beats = 0
		}
		fmt.Println(beats)
		res *= beats
	}
	// fmt.Println(input)
	return res
}
func solve2(lines []string) int {
	lines = Map(lines, func(line string, i int) string {
		return strings.ReplaceAll(line, " ", "")
	})
	return solve1(lines)
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
