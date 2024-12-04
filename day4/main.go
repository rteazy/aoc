package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne() int {
	const xmas string = "XMAS"
	total := 0

	grid := parseGrid()
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}

	var _dfs func(i, j, dx, dy, idx int)
	_dfs = func(i, j, dx, dy, idx int) {
		if idx == len(xmas)-1 {
			total += 1
			return
		}
		visited[i][j] = true
		x, y := i+dx, j+dy
		if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && string(xmas[idx+1]) == grid[x][y] {
			_dfs(x, y, dx, dy, idx+1)
		}
		visited[i][j] = false
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == "X" {
				_dfs(i, j, -1, 0, 0)
				_dfs(i, j, 1, 0, 0)
				_dfs(i, j, 0, -1, 0)
				_dfs(i, j, 0, 1, 0)
				_dfs(i, j, -1, -1, 0)
				_dfs(i, j, 1, 1, 0)
				_dfs(i, j, -1, 1, 0)
				_dfs(i, j, 1, -1, 0)
			}
		}
	}
	return total
}

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}

func partTwo() int {
	total := 0

	grid := parseGrid()
	m, n := len(grid), len(grid[0])

	isXmas := func(i, j int) bool {
		if grid[i][j] != "A" {
			return false
		}

		topLeft := NewCoordinate(i-1, j-1)
		topRight := NewCoordinate(i-1, j+1)
		bottomLeft := NewCoordinate(i+1, j-1)
		bottomRight := NewCoordinate(i+1, j+1)

		points := []Coordinate{topLeft, topRight, bottomLeft, bottomRight}
		pair := make(map[string]string)
		pair["S"] = "M"
		pair["M"] = "S"

		// checks out of bounds and that characters exists
		for _, point := range points {
			x, y := point.X, point.Y
			if x < 0 || x >= m || y < 0 || y >= n {
				return false
			}
			if _, exists := pair[grid[x][y]]; !exists {
				return false
			}
		}

		topLeftCharacter := grid[topLeft.X][topLeft.Y]
		bottomRightCharacter := grid[bottomRight.X][bottomRight.Y]
		topRightCharacter := grid[topRight.X][topRight.Y]
		bottomLeftCharacter := grid[bottomLeft.X][bottomLeft.Y]

		return pair[topLeftCharacter] == bottomRightCharacter && pair[topRightCharacter] == bottomLeftCharacter
	}

	for i := range m {
		for j := range n {
			if isXmas(i, j) {
				total += 1
			}
		}
	}
	return total
}

func parseGrid() [][]string {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
