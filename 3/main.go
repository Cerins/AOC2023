package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"os"
)

func in(x int, y int, width int, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func validSur(c byte) bool {
	return !(c == '.' || IsDigit(c))
}

func checkSur(lines []string, ni int, nj int) bool {
	width := len(lines)
	height := len(lines[0])
	if in(nj, ni, width, height) && validSur(lines[ni][nj]) {
		return true
	}
	return false
}

func surroundedByGood(lines []string, i int, j int, length int) bool {
	dl := length
	good := false
	for dl > 0 {
		// Can look on the left side of the element
		if dl == length {
			good = good || checkSur(lines, i, j-1)
			good = good || checkSur(lines, i-1, j-1)
			good = good || checkSur(lines, i+1, j-1)
		}
		// Check top
		good = good || checkSur(lines, i-1, j)
		// Check bottom
		good = good || checkSur(lines, i+1, j)
		// Can look on the right side of the element
		if dl == 1 {
			good = good || checkSur(lines, i, j+1)
			good = good || checkSur(lines, i-1, j+1)
			good = good || checkSur(lines, i+1, j+1)
		}
		j++
		dl--
	}
	return good
}

func registerGear(lines []string, ni int, nj int, number int, gears map[string][]int) {
	width := len(lines)
	height := len(lines[0])
	if in(nj, ni, width, height) && lines[ni][nj] == '*' {
		key := fmt.Sprint(ni) + ":" + fmt.Sprint(nj)
		// fmt.Println("Hello!")
		slice := gears[key]
		slice = append(slice, number)
		gears[key] = slice
	}
}

func registerGears(lines []string, i int, j int, length int, number int, gears map[string][]int) {
	dl := length
	for dl > 0 {
		// Can look on the left side of the element
		if dl == length {
			registerGear(lines, i, j-1, number, gears)
			registerGear(lines, i-1, j-1, number, gears)
			registerGear(lines, i+1, j-1, number, gears)
		}
		// Check top
		registerGear(lines, i-1, j, number, gears)
		// Check bottom
		registerGear(lines, i+1, j, number, gears)
		// Can look on the right side of the element
		if dl == 1 {
			registerGear(lines, i, j+1, number, gears)
			registerGear(lines, i-1, j+1, number, gears)
			registerGear(lines, i+1, j+1, number, gears)
		}
		j++
		dl--
	}
}

func findNumber(line string) (number int, length int) {
	negative := 1
	i := 0
	if len(line) == 0 {
		return
	}
	// Maybe negatives dont count?
	// if line[i] == '-' {
	// 	negative = -1
	// 	i++
	// }

	for ; i < len(line); i++ {
		c := line[i]
		if !IsDigit(c) {
			break
		}
		length++
		number = number*10 + ToDigit(c)
	}
	if length > 0 && negative == -1 {
		length++
	}
	number *= negative
	return
}

func solve2(lines []string) int {
	height := len(lines)
	width := len(lines[0])
	fmt.Println(lines, height, width)
	var sum int
	gears := make(map[string][]int)
	for i, line := range lines {
		numberLength := 0
		for j, _ := range line {
			if numberLength > 0 {
				numberLength--
				continue
			}
			number, length := findNumber(line[j:])
			numberLength = length
			if length > 0 {
				fmt.Println("Found a number", number, "at", j, i, "of length", length)
				registerGears(lines, i, j, length, number, gears)
			}

		}
	}
	fmt.Println(gears)
	for _, value := range gears {
		if len(value) < 2 {
			continue
		}
		res := 1
		for _, v := range value {
			res *= v
		}
		sum += res
	}
	return sum
}

func solve1(lines []string) int {
	height := len(lines)
	width := len(lines[0])
	fmt.Println(lines, height, width)
	var sum int
	for i, line := range lines {
		numberLength := 0
		for j, _ := range line {
			if numberLength > 0 {
				numberLength--
				continue
			}
			number, length := findNumber(line[j:])
			numberLength = length
			if length > 0 {
				fmt.Println("Found a number", number, "at", j, i, "of length", length)
				good := surroundedByGood(lines, i, j, length)
				fmt.Println("At it is ", good)
				if good {
					sum += number
				}
			}

		}
	}
	return sum
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
	sum := solve2(lines)
	fmt.Println(sum)
}
