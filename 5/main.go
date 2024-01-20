package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	from int
	dx   int
	// Inclusive
	till int
}

type input struct {
	seeds                 []int
	seedToSoil            []mapping
	soilToFertilizer      []mapping
	fertilizerToWater     []mapping
	waterToLight          []mapping
	lightToTemperature    []mapping
	temperatureToHumidity []mapping
	humidityToLocation    []mapping
}

func parseMapping(lines []string) ([]mapping, []string) {
	cMapping := make([]mapping, 0)
	li := 0
	for i, l := range lines {
		li = i + 1
		if i == 0 {
			continue
		}
		if l == "" {
			break
		}
		nr := strings.Split(l, " ")
		src, _ := strconv.Atoi(nr[1])
		dest, _ := strconv.Atoi(nr[0])
		rang, _ := strconv.Atoi(nr[2])
		m := mapping{
			from: src,
			dx:   dest - src,
			till: src + rang - 1,
		}
		cMapping = append(cMapping, m)
	}
	return cMapping, lines[li:]
}

func parseLines(lines []string) input {
	var input input = input{
		seeds:                 make([]int, 0),
		seedToSoil:            make([]mapping, 0),
		soilToFertilizer:      make([]mapping, 0),
		fertilizerToWater:     make([]mapping, 0),
		waterToLight:          make([]mapping, 0),
		lightToTemperature:    make([]mapping, 0),
		temperatureToHumidity: make([]mapping, 0),
		humidityToLocation:    make([]mapping, 0),
	}
	// The first line extract seeds
	for _, s := range strings.Split(strings.Split(lines[0], ": ")[1], " ") {
		n, _ := strconv.Atoi(s)
		input.seeds = append(input.seeds, n)
	}
	lines = lines[2:]
	input.seedToSoil, lines = parseMapping(lines)
	input.soilToFertilizer, lines = parseMapping(lines)
	input.fertilizerToWater, lines = parseMapping(lines)
	input.waterToLight, lines = parseMapping(lines)
	input.lightToTemperature, lines = parseMapping(lines)
	input.temperatureToHumidity, lines = parseMapping(lines)
	input.humidityToLocation, lines = parseMapping(lines)
	// The
	return input
}

func mappingsToNumber(mapping []mapping, number int) int {
	res := number
	for _, m := range mapping {
		if m.from <= number && number <= m.till {
			// fmt.Println("Transforming", res, m)
			res = number + m.dx
		}
	}
	return res
}

func seedVal(seed int, mappings [][]mapping) int {

	val := seed
	// fmt.Println(i, -1, val)
	for _, m := range mappings {
		val = mappingsToNumber(m, val)
		// fmt.Println(i, j, val)
	}
	return val
}

func solve1(lines []string) int {
	input := parseLines(lines)
	// fmt.Println(input)
	mappings := [][]mapping{
		input.seedToSoil,
		input.soilToFertilizer,
		input.fertilizerToWater,
		input.waterToLight,
		input.lightToTemperature,
		input.temperatureToHumidity,
		input.humidityToLocation,
	}
	smallest := math.MaxInt
	for _, seed := range input.seeds {
		val := seedVal(seed, mappings)
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}
func solve2(lines []string) int {
	input := parseLines(lines)
	// fmt.Println(input)
	mappings := [][]mapping{
		input.seedToSoil,
		input.soilToFertilizer,
		input.fertilizerToWater,
		input.waterToLight,
		input.lightToTemperature,
		input.temperatureToHumidity,
		input.humidityToLocation,
	}
	smallest := math.MaxInt
	for i := 0; i < len(input.seeds); i += 2 {
		sseed := input.seeds[i]
		rang := input.seeds[i+1]
		for j := 0; j < rang; j++ {
			seed := sseed + j
			val := seedVal(seed, mappings)
			if val < smallest {
				smallest = val
			}
		}
	}
	return smallest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}
	// fmt.Println(lines)
	input := solve2(lines)
	fmt.Println(input)
}
