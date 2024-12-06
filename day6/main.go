package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne() int {
	total := 0
	grid, start_x, start_y := parseInput()
	simulate(grid, start_x, start_y)
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == "X" {
				total += 1
			}
		}
	}
	return total
}

func simulate(grid [][]string, x_start, y_start int) {
	m, n := len(grid), len(grid[0])
	i, j := x_start, y_start
	dirIndex := 0
	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	for i >= 0 && i < m && j >= 0 && j < n {
		grid[i][j] = "X"
		dx, dy := directions[dirIndex][0], directions[dirIndex][1]
		x := i + dx
		y := j + dy
		if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == "#" {
			dirIndex = (dirIndex + 1) % 4
			continue
		}
		i, j = x, y
	}
}

func parseInput() ([][]string, int, int) {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startX, startY := -1, -1

	grid := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, "")
		grid = append(grid, characters)
	}

	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == "^" {
				startX = i
				startY = j
			}
		}
	}

	return grid, startX, startY
}

func partTwo() int {
	total := 0
	grid, start_x, start_y := parseInput()
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == "." {
				grid[i][j] = "#"
				if infinitePath(grid, start_x, start_y) {
					total += 1
				}
				grid[i][j] = "."
			}
		}
	}
	return total
}

type Point struct {
	X, Y, DX, DY int
}

func NewPoint(x, y, dx, dy int) Point {
	return Point{x, y, dx, dy}
}

func infinitePath(grid [][]string, startX, startY int) bool {
	seen := make(map[Point]bool)

	m, n := len(grid), len(grid[0])
	i, j := startX, startY
	dirIndex := 0
	directions := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	for i >= 0 && i < m && j >= 0 && j < n {
		dx, dy := directions[dirIndex][0], directions[dirIndex][1]
		point := NewPoint(i, j, dx, dy)
		if _, exists := seen[point]; exists {
			return true
		} else {
			seen[point] = true
		}

		x := i + dx
		y := j + dy
		if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == "#" {
			dirIndex = (dirIndex + 1) % 4
			continue
		}
		i, j = x, y
	}

	return false
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
