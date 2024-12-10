package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() int {
	res := 0
	grid := parseInput()
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 0 {
				res += score(grid, i, j)
			}
		}
	}

	return res
}

func parseInput() [][]int {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		vals := []int{}
		for _, c := range line {
			var val int
			if c == '.' {
				val = -1
			} else {
				val, err = strconv.Atoi(string(c))
				if err != nil {
					log.Fatal(err)
				}
			}
			vals = append(vals, val)

		}
		grid = append(grid, vals)
	}

	return grid
}

func score(grid [][]int, startI, startJ int) int {
	start := []int{startI, startJ}
	count := 0
	queue := [][]int{start}
	visited := [][]bool{}
	for range len(grid) {
		arr := []bool{}
		for range len(grid[0]) {
			arr = append(arr, false)
		}
		visited = append(visited, arr)

	}
	visited[startI][startJ] = true

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		if grid[i][j] == 9 {
			count += 1
			continue
		}

		dir := [][]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1},
		}

		for _, point := range dir {
			x, y := point[0], point[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == grid[i][j]+1 && !visited[x][y] {
				visited[x][y] = true
				queue = append(queue, point)
			}
		}

	}

	return count
}

func scoreUsingAllPaths(grid [][]int, startI, startJ int) int {
	start := []int{startI, startJ}
	count := 0
	queue := [][]int{start}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		if grid[i][j] == 9 {
			count += 1
			continue
		}

		dir := [][]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1},
		}

		for _, point := range dir {
			x, y := point[0], point[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == grid[i][j]+1 {
				queue = append(queue, point)
			}
		}

	}

	return count
}

func partTwo() int {
	res := 0
	grid := parseInput()
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 0 {
				res += scoreUsingAllPaths(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
