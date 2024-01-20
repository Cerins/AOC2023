package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func magicValue(color string) int {
	switch color {
	case "blue":
		return 14
	case "green":
		return 13
	case "red":
		return 12
	default:
		return 0
	}
}

func solveLine1(txt string) int {
	txt = strings.ReplaceAll(txt, "; ", ", ")
	fmt.Println(txt)
	spl := strings.Split(txt, ": ")
	start := spl[0]
	values := spl[1]
	valid := 1
	game, _ := strconv.Atoi(strings.Split(start, " ")[1])
	fmt.Println("Game", game)
	for _, v := range strings.Split(values, ", ") {
		spl := strings.Split(v, " ")
		count, _ := strconv.Atoi(spl[0])
		color := spl[1]
		mgik := magicValue(color)
		fmt.Println("Color", color, "count", count, "mgik", mgik)
		if count > magicValue(spl[1]) {
			valid = 0
			break
		}
	}
	return game * valid
}

func solveLine2(txt string) int {
	txt = strings.ReplaceAll(txt, "; ", ", ")
	fmt.Println(txt)
	spl := strings.Split(txt, ": ")
	start := spl[0]
	values := spl[1]
	game, _ := strconv.Atoi(strings.Split(start, " ")[1])
	var mr, mg, mb int
	fmt.Println("Game", game)
	for _, v := range strings.Split(values, ", ") {
		spl := strings.Split(v, " ")
		count, _ := strconv.Atoi(spl[0])
		color := spl[1]
		fmt.Println("Color", color, "count", count)
		switch color {
		case "red":
			mr = Max(mr, count)
		case "blue":
			mb = Max(mb, count)
		case "green":
			mg = Max(mg, count)
		}
	}
	return mr * mb * mg

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		txt := scanner.Text()
		sum += solveLine2(txt)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
