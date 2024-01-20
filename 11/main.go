package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"os"
)

type Cord struct {
	x int
	y int
}

func dist(a *Cord, b *Cord) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

func galaxyCords(lines []string) []Cord {
	cords := make([]Cord, 0)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				cords = append(cords, Cord{
					x,
					y,
				})
			}
		}
	}
	return cords
}
func galaxyCords2(lines []string, emptyRow, emptyColumn []bool, to int) []Cord {
	cords := make([]Cord, 0)
	skippedY := 0
	for y, line := range lines {
		if emptyRow[y] {
			skippedY++
		}
		skippedX := 0
		for x, c := range line {
			if emptyColumn[x] {
				skippedX++
			}
			if c == '#' {
				cords = append(cords, Cord{
					x + skippedX*(to-1),
					y + skippedY*(to-1),
				})
			}
		}
	}
	return cords
}

func prettyPrint(lines []string) {
	for _, line := range lines {
		for _, c := range line {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}
func expandRules(lines []string) ([]bool, []bool) {
	emptyRow := make([]bool, len(lines))
	emptyColumn := make([]bool, len(lines[0]))
	// Find all the empty rows
	for y := 0; y < len(lines); y++ {
		empty := true
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] != '.' {
				empty = false
			}
		}
		emptyRow[y] = empty
	}
	for x := 0; x < len(lines[0]); x++ {
		empty := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != '.' {
				empty = false
			}
		}
		emptyColumn[x] = empty
	}
	return emptyRow, emptyColumn
}

func expand(lines []string, to int) []string {
	nLines := make([]string, 0)
	emptyRow, emptyColumn := expandRules(lines)
	for y, line := range lines {
		row := ""
		for x, c := range line {
			row += string(c)
			if emptyColumn[x] {
				for z := 1; z < to; z++ {

					row += "."
				}
			}
		}
		// fmt.Println(row)
		nLines = append(nLines, row)
		if emptyRow[y] {
			for z := 1; z < to; z++ {
				nLines = append(nLines, row)
			}
		}
	}
	// fmt.Println(emptyColumn)
	// fmt.Println(emptyRow)
	// prettyPrint(nLines)
	return nLines
}

func solve1(lines []string) int {
	// fmt.Println((lines))
	// lines = expand(lines, 2)
	emptyRow, emptyColumn := expandRules(lines)
	cords := galaxyCords2(lines, emptyRow, emptyColumn, 2)
	res := 0
	for len(cords) > 0 {
		c1 := cords[0]
		cords = cords[1:]
		for _, c2 := range cords {
			res += dist(&c1, &c2)
		}
	}
	return res
}
func solve2(lines []string) int {
	emptyRow, emptyColumn := expandRules(lines)
	cords := galaxyCords2(lines, emptyRow, emptyColumn, 1000000)
	res := 0
	// Funnily enough this actually does not bring any optimization benefits
	for k := 0; k < len(cords); k++ {
		c1 := cords[k]
		// This is actually fast
		scords := cords[k:]
		for _, c2 := range scords {
			res += dist(&c1, &c2)
		}
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
	solution := solve2(lines)
	fmt.Println(solution)
}
