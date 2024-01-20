package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func solveLine1(line string) int {
	first, last := -1, -1
	for _, char := range line {
		isdigit := unicode.IsDigit(char)
		if isdigit {
			digit := int(char - '0')
			if first == -1 {
				first = digit
			}
			last = digit
		}
	}
	return first*10 + last
}

func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

type number struct {
	spelling string
	value    int
}

var spellings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}
var options = Map(spellings, func(v string, _ int) number {
	value := ((FirstOcc(spellings, v) + 1) % 10)
	// fmt.Println(v, value)
	return number{spelling: v, value: value}
})

func solveLine2(line string) int {
	var nline string
	index := 0
	for {
		newIndex := -1
		cO := 0
		for newCO, op := range options {
			tmpIndex := indexAt(line, op.spelling, index)
			if tmpIndex != -1 && (newIndex == -1 || tmpIndex < newIndex) {
				cO = newCO
				newIndex = tmpIndex
			}
		}
		if newIndex == -1 {
			break
		}

		index = newIndex + 1
		nline += fmt.Sprint(options[cO].value)
	}
	// fmt.Println(nline)
	return solveLine1(nline)
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
