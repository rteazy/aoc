package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne() int {
	res := 0
	grid, moves, position := parseInput()
	m, n := len(grid), len(grid[0])

	var moveRobot func(position Point, move string) (Point, bool)
	moveRobot = func(position Point, move string) (Point, bool) {
		i, j := position.X, position.Y
		dx, dy := 0, 0
		switch move {
		case "^":
			dx -= 1
		case ">":
			dy += 1
		case "v":
			dx += 1
		case "<":
			dy -= 1
		default:
			log.Fatalf("Move invalid: %s\n", move)
		}

		x, y := position.X+dx, position.Y+dy
		if grid[x][y] == "#" {
			return position, false
		}

		nextPosition := NewPoint(x, y)
		if grid[x][y] == "O" {
			if _, shifted := moveRobot(nextPosition, move); !shifted {
				return position, false
			}
		}

		grid[x][y] = grid[i][j]
		grid[i][j] = "."
		return nextPosition, true
	}

	for _, move := range moves {
		position, _ = moveRobot(position, move)
	}
	for i := range m {
		for j := range n {
			if grid[i][j] == "O" {
				res += 100*i + j
			}
		}
	}

	return res
}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func parseInput() ([][]string, []string, Point) {
	grid := [][]string{}
	moves := []string{}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		grid = append(grid, strings.Split(line, ""))
	}
	for scanner.Scan() {
		for _, m := range scanner.Text() {
			moves = append(moves, string(m))
		}
	}

	var start Point
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == "@" {
				start = NewPoint(i, j)
			}
		}
	}

	return grid, moves, start
}

func partTwo() int {
	res := 0
	firstGrid, moves, _ := parseInput()
	grid := alterGrid(firstGrid)
	m, n := len(grid), len(grid[0])
	position := getPosition(grid)

	var moveRobot func(position Point, move string, simulateMove bool) (Point, bool)
	moveRobot = func(position Point, move string, simulateMove bool) (Point, bool) {
		dx, dy := 0, 0
		switch move {
		case "^":
			dx -= 1
		case ">":
			dy += 1
		case "v":
			dx += 1
		case "<":
			dy -= 1
		default:
			log.Fatalf("Move invalid: %s\n", move)
		}

		x, y := position.X+dx, position.Y+dy

		if grid[x][y] == "#" {
			return position, false
		}

		nextPosition := NewPoint(x, y)

		// box found
		if grid[x][y] == "[" || grid[x][y] == "]" {
			if move == "<" || move == ">" {
				if _, shifted := moveRobot(nextPosition, move, false); !shifted {
					return position, false
				}
			} else {
				if _, moveForwardSim := moveRobot(nextPosition, move, true); !moveForwardSim {
					return position, false
				}
				complement := NewPoint(-1, -1)
				if grid[x][y] == "[" {
					complement = NewPoint(x, y+1)
					if _, rightMoveSim := moveRobot(complement, move, true); !rightMoveSim {
						return position, false
					}
				} else {
					complement = NewPoint(x, y-1)
					if _, leftMoveSim := moveRobot(complement, move, true); !leftMoveSim {
						return position, false
					}
				}

				moveRobot(nextPosition, move, simulateMove)
				moveRobot(complement, move, simulateMove)
			}
		}

		if simulateMove {
			return nextPosition, true
		}
		grid[x][y] = grid[position.X][position.Y]
		grid[position.X][position.Y] = "."

		return nextPosition, true
	}

	for _, move := range moves {
		position, _ = moveRobot(position, move, false)
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == "[" {
				res += 100*i + j
			}
		}
	}

	return res
}

func printGrid(grid [][]string) {
	for i := range len(grid) {
		for j := range len(grid[0]) {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func alterGrid(original [][]string) [][]string {
	grid := [][]string{}
	m, n := len(original), len(original[0])
	for i := range m {
		row := []string{}
		for j := range n {
			if original[i][j] == "#" {
				row = append(row, "#")
				row = append(row, "#")
			} else if original[i][j] == "O" {
				row = append(row, "[")
				row = append(row, "]")
			} else if original[i][j] == "." {
				row = append(row, ".")
				row = append(row, ".")
			} else {
				row = append(row, "@")
				row = append(row, ".")
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func getPosition(grid [][]string) Point {
	var position Point
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == "@" {
				position = NewPoint(i, j)
			}
		}
	}
	return position
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
