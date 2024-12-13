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
	return res
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
