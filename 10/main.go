package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cord struct {
	x int
	y int
}

type Move struct {
	to   Cord
	from *Cord
	d    int
}

func (c *Cord) add(o *Cord) Cord {
	var n Cord
	n.x = c.x + o.x
	n.y = c.y + o.y
	return n
}

func (c *Cord) valid(width, height int) bool {
	return c.x >= 0 && c.x < width && c.y >= 0 && c.y < height
}

func nextMoves(lines []string, m *Move) ([]Move, bool) {
	moves := make([]Move, 0)
	valid := true
	x := m.to.x
	y := m.to.y

	trigger := lines[y][x]

	switch trigger {
	case 'S':
		{
			// Add all the 4 passings no questions asked
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: -1,
					y: 0,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 1,
					y: 0,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 0,
					y: 1,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 0,
					y: -1,
				}),
			})
		}
	case '-':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fx := from.x
				if x-fx == 1 || x-fx == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: x - fx,
							y: 0,
						}),
					})
				}
			}

		}
	case '|':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				if y-fy == 1 || y-fy == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: y - fy,
						}),
					})
				}
			}

		}
	case 'J':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				fx := from.x
				if x-fx == 1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: -1,
						}),
					})
				} else if y-fy == 1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: -1,
							y: 0,
						}),
					})
				}
			}

		}
	case '7':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				fx := from.x
				if x-fx == 1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: 1,
						}),
					})
				} else if y-fy == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: -1,
							y: 0,
						}),
					})
				}
			}

		}
	case 'L':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				fx := from.x
				if x-fx == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: -1,
						}),
					})
				} else if y-fy == 1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 1,
							y: 0,
						}),
					})
				}
			}

		}
	case 'F':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				fx := from.x
				if x-fx == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: 1,
						}),
					})
				} else if y-fy == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 1,
							y: 0,
						}),
					})
				}
			}

		}
	}

	return moves, valid
}

func nextMoves2(lines []string, m *Move) ([]Move, bool) {
	moves := make([]Move, 0)
	valid := true
	x := m.to.x
	y := m.to.y

	trigger := lines[y][x]

	switch trigger {
	case 'S':
		{
			// Add all the 4 passings no questions asked
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: -1,
					y: 0,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 1,
					y: 0,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 0,
					y: 1,
				}),
			})
			moves = append(moves, Move{
				from: &m.to,
				d:    m.d + 1,
				to: m.to.add(&Cord{
					x: 0,
					y: -1,
				}),
			})
		}
	case '-':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fx := from.x
				if x-fx == 1 || x-fx == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: x - fx,
							y: 0,
						}),
					})
				}
			}

		}
	case '|':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				if y-fy == 1 || y-fy == -1 {
					valid = true
					moves = append(moves, Move{
						from: &m.to,
						d:    m.d + 1,
						to: m.to.add(&Cord{
							x: 0,
							y: y - fy,
						}),
					})
				}
			}

		}
	case 'J', 'L', '7', 'F':
		{
			// Assume that invalid
			valid = false
			from := m.from
			if from != nil {
				fy := from.y
				fx := from.x
				valid = true
				moves = append(moves, Move{
					from: &m.to,
					d:    m.d + 1,
					to: m.to.add(&Cord{
						x: x - fx,
						y: y - fy,
					}),
				})
			}

		}
	}

	return moves, valid
}

func printMaze(lines []string, d [][]int) {
	for y, row := range d {
		for x, item := range row {
			if lines[y][x] == 'S' {
				fmt.Print("S")
			} else if lines[y][x] == '.' {
				fmt.Print(".")
			} else if item >= 10 {
				fmt.Print("X")
			} else {
				fmt.Print(item)
			}
		}
		fmt.Println()
	}
}
func printVisits(d [][]bool) {
	for _, row := range d {
		for _, item := range row {
			if item {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func solve1(lines []string) (int, [][]bool, [][]int) {
	height := len(lines)
	width := len(lines[0])
	visited := make([][]bool, height)
	distance := make([][]int, height)
	for i := range visited {
		visited[i] = make([]bool, width)
		distance[i] = make([]int, width)
	}
	que := make([]Move, 0)
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				que = append(que, Move{
					to: Cord{
						x,
						y,
					},
					d: 1,
				})
				// Goto supremacy
				goto start
			}
		}
	}
start:
	for len(que) > 0 {
		// First pop the cord
		item := que[0]
		que = que[1:]
		// Check if cord is valid
		if !item.to.valid(width, height) {
			continue
		}
		x := item.to.x
		y := item.to.y
		if visited[y][x] == true {
			continue
		}
		next, goodVisit := nextMoves(lines, &item)
		que = append(que, next...)
		if goodVisit {
			visited[y][x] = true
			if lines[y][x] != '.' {
				distance[y][x] = item.d
			}
		}
	}
	// printMaze(lines, distance)
	res := 0
	for _, row := range distance {
		for _, item := range row {
			if item > res {
				res = item
			}
		}
	}
	return res - 1, visited, distance
}

func solve2(lines []string) int {
	height := len(lines)
	width := len(lines[0])
	_, _, distances := solve1(lines)
	// Create a maze 3 times the original
	visited := make([][]bool, height*3)
	for i := range visited {
		visited[i] = make([]bool, width*3)
	}
	// Make some ground rules
	for y, row := range distances {
		for x, d := range row {
			// If visited then the new grid is also set
			visited[y*3+1][x*3+1] = d > 0
			if d > 0 {
				from := Cord{
					x,
					y,
				}
				var to Cord
				// Check left
				to.x = x - 1
				to.y = y
				var mark bool
				if to.valid(width, height) {
					_, mark = nextMoves(lines, &Move{
						to,
						&from,
						d,
					})
					visited[y*3+1][x*3] = distances[to.y][to.x] > 0 && mark
				}
				// Check right
				to.x = x + 1
				to.y = y
				if to.valid(width, height) {
					_, mark = nextMoves(lines, &Move{
						to,
						&from,
						d,
					})
					visited[y*3+1][x*3+2] = distances[to.y][to.x] > 0 && mark
				}
				// Check top
				to.x = x
				to.y = y - 1
				if to.valid(width, height) {
					_, mark = nextMoves(lines, &Move{
						to,
						&from,
						d,
					})
					visited[y*3][x*3+1] = distances[to.y][to.x] > 0 && mark
				}
				// Check bottom
				to.x = x
				to.y = y + 1
				if to.valid(width, height) {
					_, mark = nextMoves(lines, &Move{
						to,
						&from,
						d,
					})
					visited[y*3+2][x*3+1] = distances[to.y][to.x] > 0 && mark
				}
			}

		}
	}
	// printVisits(visited)
	que := make([]Move, 0)
	// Then put every corner as a potential escape route
	for y, line := range visited {
		for x, _ := range line {
			if x == 0 {
				que = append(que, Move{
					to: Cord{
						x,
						y,
					},
					d: 0,
				})
			}
			if x == len(line)-1 {
				que = append(que, Move{
					to: Cord{
						x,
						y,
					},
					d: 0,
				})
			}
			if y == 0 {
				que = append(que, Move{
					to: Cord{
						x,
						y,
					},
					d: 0,
				})
			}
			if y == len(visited)-1 {
				que = append(que, Move{
					to: Cord{
						x,
						y,
					},
					d: 0,
				})
			}
		}
	}

	for len(que) > 0 {
		item := que[0]
		que = que[1:]
		x := item.to.x
		y := item.to.y
		// Skip invalid
		if !item.to.valid(width*3, height*3) {
			continue
		}
		// Do nothing if visited
		if visited[y][x] {
			continue
		}
		visited[y][x] = true
		// Try to visit all neighbors
		que = append(que, Move{
			to: item.to.add(&Cord{
				x: 1,
				y: 0,
			}),
			from: &item.to,
			d:    0,
		},
		)
		que = append(que, Move{
			to: item.to.add(&Cord{
				x: -1,
				y: 0,
			}),
			from: &item.to,
			d:    0,
		},
		)
		que = append(que, Move{
			to: item.to.add(&Cord{
				x: 0,
				y: -1,
			}),
			from: &item.to,
			d:    0,
		},
		)
		que = append(que, Move{
			to: item.to.add(&Cord{
				x: 0,
				y: 1,
			}),
			from: &item.to,
			d:    0,
		},
		)
	}
	// printVisits(visited)
	res := 0
	// Now that all remains is to see which points where unvisited
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !visited[y*3+1][x*3+1] {
				res++
			}
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
