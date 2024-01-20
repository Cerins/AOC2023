package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type node struct {
	left  string
	right string
}

type game struct {
	start    string
	end      string
	sequence string
	lookup   map[string]node
}

func parseLines(lines []string) game {
	var game game
	game.sequence = lines[0]
	game.start = "AAA"
	game.end = "ZZZ"
	game.lookup = make(map[string]node)
	r := regexp.MustCompile(`^([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)$`)
	for _, line := range lines[2:] {
		match := r.FindStringSubmatch(line)
		key := match[1]

		game.lookup[key] = node{
			left:  match[2],
			right: match[3],
		}
		// fmt.Printf("%#v\n", r.FindStringSubmatch(line))
	}
	return game
}

func solve1(lines []string) int {
	res := 0
	game := parseLines(lines)
	c := game.start
	i := 0
	for c != game.end {
		command := game.sequence[i]
		i = (i + 1) % len(game.sequence)
		switch command {
		case 'L':
			c = game.lookup[c].left
		case 'R':
			c = game.lookup[c].right
		}
		res++
	}
	// fmt.Println(game)
	return res
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func solve2LCM(lines []string) int64 {
	game := parseLines(lines)
	c := make([]string, 0)
	// Go through
	i := 0
	for k := range game.lookup {
		lastChar := k[len(k)-1]
		if lastChar == 'A' {
			c = append(c, k)
		}
	}
	pathC := make([]int64, len(c))
	for j, k := range c {
		var path int64
		for {
			lastChar := k[len(k)-1]
			if lastChar == 'Z' {
				break
			}
			command := game.sequence[i]
			i = (i + 1) % len(game.sequence)
			switch command {
			case 'L':
				k = game.lookup[k].left
			case 'R':
				k = game.lookup[k].right
			}
			path++
		}
		pathC[j] = path
	}
	// fmt.Println(pathC)
	return LCM(pathC[0], pathC[1], pathC[2:]...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}
	// fmt.Println(lines)
	res := solve2LCM(lines)
	fmt.Println(res)
}
