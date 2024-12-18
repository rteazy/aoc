package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func partOne() int {
	q1Count, q2Count, q3Count, q4Count := 0, 0, 0, 0
	m, n := 101, 103

	robots := parseInput()
	positions := []Point{}
	for _, robot := range robots {
		positions = append(positions, move(robot, 100, m, n))
	}

	for _, pos := range positions {
		if pos.X < m/2 && pos.Y < n/2 {
			q1Count += 1
		} else if pos.X > m/2 && pos.Y < n/2 {
			q2Count += 1
		} else if pos.X < m/2 && pos.Y > n/2 {
			q3Count += 1
		} else if pos.X > m/2 && pos.Y > n/2 {
			q4Count += 1
		}
	}

	return q1Count * q2Count * q3Count * q4Count
}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type Robot struct {
	Position Point
	Velocity Point
}

func NewRobot(pos, vel Point) Robot {
	return Robot{pos, vel}
}

func parseInput() []Robot {
	robots := []Robot{}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("p=(\\d+),(\\d+) v=(\\-?\\d+),(\\-?\\d+)")
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)[1:]
		posX, posY := matches[0], matches[1]
		velX, velY := matches[2], matches[3]
		x, err := strconv.Atoi(posX)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(posY)
		if err != nil {
			log.Fatal(err)
		}
		vx, err := strconv.Atoi(velX)
		if err != nil {
			log.Fatal(err)
		}
		vy, err := strconv.Atoi(velY)
		if err != nil {
			log.Fatal(err)
		}
		pos := NewPoint(x, y)
		vel := NewPoint(vx, vy)
		robots = append(robots, NewRobot(pos, vel))
	}

	return robots
}

func move(robot Robot, seconds, m, n int) Point {
	lastPos := robot.Position
	for range seconds {
		x := lastPos.X + robot.Velocity.X
		if x < 0 {
			x = m + x
		} else if x >= m {
			x = x % m
		}
		y := lastPos.Y + robot.Velocity.Y
		if y < 0 {
			y = n + y
		} else if y >= n {
			y = y % n
		}
		lastPos = NewPoint(x, y)
	}
	return lastPos
}

func partTwo() int {
	m, n := 101, 103
	robots := parseInput()
	positions := []Point{}
	velocities := []Point{}

	for _, robot := range robots {
		positions = append(positions, robot.Position)
		velocities = append(velocities, robot.Velocity)
	}

	res := 0
	for second := range 10000 {
		// assumption (a bad one) - no overlapping robots will have the XMAS tree
		positions = moveAllRobotsOnce(positions, velocities, m, n)
		seen := make(map[Point]bool)
		for _, p := range positions {
			if _, exists := seen[p]; !exists {
				seen[p] = true
			}
		}
		if len(seen) == len(robots) {
			printGrid(positions, m, n)
			res = second + 1
		}
	}

	return res
}

func moveAllRobotsOnce(positions, velocities []Point, m, n int) []Point {
	newPositions := []Point{}
	for i, pos := range positions {
		x := pos.X + velocities[i].X
		if x < 0 {
			x = m + x
		} else if x >= m {
			x = x % m
		}
		y := pos.Y + velocities[i].Y
		if y < 0 {
			y = n + y
		} else if y >= n {
			y = y % n
		}
		newPositions = append(newPositions, NewPoint(x, y))
	}
	return newPositions
}

func printGrid(points []Point, m, n int) {
	grid := [][]string{}
	for range n {
		r := []string{}
		for range m {
			r = append(r, ".")
		}
		grid = append(grid, r)
	}
	for _, pos := range points {
		grid[pos.Y][pos.X] = "*"
	}

	f, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for y := range n {
		for x := range m {
			f.WriteString(grid[y][x])
		}
		f.WriteString("\n")
	}

	f.WriteString("\n")
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
