package main

import (
	. "AOC2023/utils"
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Puzzle struct {
	line string
	size []int
	max  int
}

func lineToPuzzle(line string, _ int) Puzzle {
	parsed := strings.Split(line, " ")
	var p Puzzle
	p.line = parsed[0]
	strN := strings.Split(parsed[1], ",")
	p.size = make([]int, len(strN))
	for i, sN := range strN {
		p.size[i], _ = strconv.Atoi(sN)
		p.max += p.size[i]
	}
	return p
}

func lineToPuzzlePart2(line string, _ int) Puzzle {
	N := 5
	parsed := strings.Split(line, " ")
	var p Puzzle
	p.line = strings.Repeat(parsed[0]+"?", N)
	p.line = p.line[:len(p.line)-1] // Remove the last '?'

	strN := strings.Split(parsed[1], ",")
	newSizes := make([]string, 0, len(strN)*N)
	for i := 0; i < N; i++ {
		for _, sN := range strN {
			// Repeat each size 5 times
			newSizes = append(newSizes, sN)
		}
	}
	p.size = make([]int, len(newSizes))
	for i, sN := range newSizes {
		p.size[i], _ = strconv.Atoi(sN)
		p.max += p.size[i]
	}
	return p
}

type Memoizable func(Memoizable, string, []int) int

type dpKey struct {
	remainder string
	sizeLen   int
}

func hashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func Memo(fn Memoizable) Memoizable {
	mem := make(map[dpKey]int)
	var memoizedFunc Memoizable
	memoizedFunc = func(_ Memoizable, s string, i []int) int {
		key := dpKey{remainder: s, sizeLen: len(i)}
		if result, exists := mem[key]; exists {
			return result
		}
		result := fn(memoizedFunc, s, i)
		mem[key] = result
		return result
	}
	return memoizedFunc
}

func sum(a []int) int {
	s := 0
	for _, c := range a {
		s += c
	}
	return s
}

func dp(self Memoizable, remainder string, size []int) int {
	// fmt.Println(remainder, size)
	if len(remainder) == 0 {
		if len(size) == 0 {
			return 1
		}
		return 0
	}
	if len(size) == 0 {
		// Check if the remainder does has #
		has := false
		for _, c := range remainder {
			has = has || c == '#'
		}
		// If has return 0
		if has {
			return 0
		}
		return 1
	}
	// Not enough items
	if len(remainder) < sum(size)+len(size)-1 {
		return 0
	}
	c := remainder[0]
	nr := remainder[1:]
	if c == '.' {
		return self(self, nr, size)
	}
	if c == '#' {
		if len(remainder) < size[0] {
			return 0
		}
		i := 1
		for ; i < size[0]; i++ {
			if remainder[i] == '.' {
				return 0
			}
		}
		if i < len(remainder) && remainder[i] == '#' {
			return 0
		}
		if i < len(remainder) {
			i++
		}
		return self(self, remainder[i:], size[1:])
	}
	return self(self, "#"+nr, size) + self(self, "."+nr, size)
}

func solve1(lines []string) int {
	// fmt.Println(lines)
	puzzles := Map(lines, lineToPuzzle)
	// fmt.Println(puzzles)
	res := 0
	for _, p := range puzzles {
		// fmt.Println(p.line)
		// fmt.Println("---")
		dpMemo := Memo(dp)
		calc := dpMemo(dp, p.line, p.size)
		// fmt.Println("---")
		// fmt.Println(p.line, calc)
		res += calc
	}
	return res
}

func solve2(lines []string) int64 {
	puzzles := Map(lines, lineToPuzzlePart2)
	var wg sync.WaitGroup

	resCh := make(chan int, len(puzzles)) // Channel for results
	semaphore := make(chan struct{}, 64)  // Semaphore to limit goroutines

	for i, p := range puzzles {
		wg.Add(1)               // Increment the WaitGroup counter
		semaphore <- struct{}{} // Acquire a slot in the semaphore

		go func(p Puzzle, i int) {
			defer wg.Done()                // Decrement the WaitGroup counter when goroutine completes
			defer func() { <-semaphore }() // Release the slot in the semaphore

			dpMemo := Memo(dp)
			calc := dpMemo(dpMemo, p.line, p.size)
			// fmt.Println("Solved", i+1, calc)

			resCh <- calc // Send result to the channel
		}(p, i)
	}

	wg.Wait()    // Wait for all goroutines to complete
	close(resCh) // Close the channel

	var res int64
	for calc := range resCh { // Collect results from the channel
		res += int64(calc)
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
