package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partOne() int {
	res := 0

	visited := [][]bool{}
	grid := parseInput()
	for range len(grid) {
		gridEntry := []bool{}
		for range len(grid[0]) {
			gridEntry = append(gridEntry, false)
		}
		visited = append(visited, gridEntry)
	}

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if !visited[i][j] {
				res += getPrice(grid, visited, i, j)
			}
		}
	}
	return res
}

func getPrice(grid [][]string, visited [][]bool, startI, startJ int) int {
	area, perimeter := 0, 0
	queue := [][]int{
		{startI, startJ},
	}

	visited[startI][startJ] = true

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		i, j := point[0], point[1]
		area += 1
		perimeter += 4
		nextPoints := [][]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1},
		}
		for _, nextPoint := range nextPoints {
			x, y := nextPoint[0], nextPoint[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) &&
				grid[x][y] == grid[i][j] {
				perimeter -= 1
				if !visited[x][y] {
					visited[x][y] = true
					queue = append(queue, []int{x, y})
				}
			}
		}
	}

	return area * perimeter
}

func parseInput() [][]string {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	for scanner.Scan() {
		gridEntry := []string{}
		line := scanner.Text()
		for _, c := range line {
			gridEntry = append(gridEntry, string(c))
		}
		grid = append(grid, gridEntry)
	}

	return grid
}

func partTwo() int {
	res := 0

	visited := [][]bool{}
	grid := parseInput()
	for range len(grid) {
		gridEntry := []bool{}
		for range len(grid[0]) {
			gridEntry = append(gridEntry, false)
		}
		visited = append(visited, gridEntry)
	}

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if !visited[i][j] {
				res += getPriceWithSides(grid, visited, i, j)
			}
		}
	}
	return res
}

func getPriceWithSides(grid [][]string, visited [][]bool, startI, startJ int) int {
	area, sides := 0, 0
	queue := [][]int{
		{startI, startJ},
	}

	visited[startI][startJ] = true

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		i, j := point[0], point[1]
		area += 1
		sides += countCorners(grid, i, j)
		nextPoints := [][]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1},
		}
		for _, nextPoint := range nextPoints {
			x, y := nextPoint[0], nextPoint[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) &&
				grid[x][y] == grid[i][j] {
				if !visited[x][y] {
					visited[x][y] = true
					queue = append(queue, []int{x, y})
				}
			}
		}
	}

	return area * sides
}

func countCorners(grid [][]string, i, j int) int {
	corners := 0
	m, n := len(grid), len(grid[0])
	curr := grid[i][j]

	var top, bottom, left, right string
	topX, topY := i-1, j
	bottomX, bottomY := i+1, j
	leftX, leftY := i, j-1
	rightX, rightY := i, j+1
	if topX >= 0 && topX < m && topY >= 0 && topY < n {
		top = grid[topX][topY]
	} else {
		top = "."
	}
	if bottomX >= 0 && bottomX < m && bottomY >= 0 && bottomY < n {
		bottom = grid[bottomX][bottomY]
	} else {
		bottom = "."
	}
	if leftX >= 0 && leftX < m && leftY >= 0 && leftY < n {
		left = grid[leftX][leftY]
	} else {
		left = "."
	}
	if rightX >= 0 && rightX < m && rightY >= 0 && rightY < n {
		right = grid[rightX][rightY]
	} else {
		right = "."
	}

	var cornerTopLeft, cornerTopRight, cornerBottomLeft, cornerBottomRight string
	cornerTopLeftX, cornerTopLeftY := i-1, j-1
	cornerTopRightX, cornerTopRightY := i-1, j+1
	cornerBottomLeftX, cornerBottomLeftY := i+1, j-1
	cornerBottomRightX, cornerBottomRightY := i+1, j+1
	if cornerTopLeftX >= 0 && cornerTopLeftX < m && cornerTopLeftY >= 0 && cornerTopLeftY < n {
		cornerTopLeft = grid[cornerTopLeftX][cornerTopLeftY]
	} else {
		cornerTopLeft = "."
	}

	if cornerTopRightX >= 0 && cornerTopRightX < m && cornerTopRightY >= 0 && cornerTopRightY < n {
		cornerTopRight = grid[cornerTopRightX][cornerTopRightY]
	} else {
		cornerTopRight = "."
	}

	if cornerBottomLeftX >= 0 && cornerBottomLeftX < m && cornerBottomLeftY >= 0 && cornerBottomLeftY < n {
		cornerBottomLeft = grid[cornerBottomLeftX][cornerBottomLeftY]
	} else {
		cornerBottomLeft = "."
	}

	if cornerBottomRightX >= 0 && cornerBottomRightX < m && cornerBottomRightY >= 0 && cornerBottomRightY < n {
		cornerBottomRight = grid[cornerBottomRightX][cornerBottomRightY]
	} else {
		cornerBottomRight = "."
	}

	// top left corner
	if curr == top && curr == left && curr != cornerTopLeft {
		corners += 1
	} else if curr != top && curr != left {
		corners += 1
	}

	// top right corner
	if curr == top && curr == right && curr != cornerTopRight {
		corners += 1
	} else if curr != top && curr != right {
		corners += 1
	}

	// bottom left corner
	if curr == bottom && curr == left && curr != cornerBottomLeft {
		corners += 1
	} else if curr != bottom && curr != left {
		corners += 1
	}

	// bottom right corner
	if curr == bottom && curr == right && curr != cornerBottomRight {
		corners += 1
	} else if curr != bottom && curr != right {
		corners += 1
	}

	return corners
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
