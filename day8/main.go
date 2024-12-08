package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func partOne() int {
	antinodes := make(map[Point]bool)
	frequencies, m, n := parseInput()
	for _, points := range frequencies {
		findAntiNodes(antinodes, points, m, n)
	}

	return len(antinodes)

}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func parseInput() (map[string][]Point, int, int) {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	frequencies := make(map[string][]Point)
	scanner := bufio.NewScanner(file)
	i := 0
	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		n = len(line)
		chs := strings.Split(line, "")
		for j, c := range chs {
			if c != "." {
				point := NewPoint(i, j)
				frequencies[c] = append(frequencies[c], point)
			}
		}
		i += 1
	}

	return frequencies, i, n
}

func findAntiNodes(antinodes map[Point]bool, points []Point, m, n int) {
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			distX := math.Pow(float64(a.X-b.X), 2)
			distY := math.Pow(float64(a.Y-b.Y), 2)
			dist := math.Sqrt(distX + distY)

			if dist < 1 {
				continue
			}

			aAntinodeX := a.X + (a.X - b.X)
			aAntinodeY := a.Y + (a.Y - b.Y)
			if aAntinodeX >= 0 && aAntinodeX < m && aAntinodeY >= 0 && aAntinodeY < n {
				point := NewPoint(aAntinodeX, aAntinodeY)
				antinodes[point] = true
			}

			bAntinodeX := b.X + (b.X - a.X)
			bAntinodeY := b.Y + (b.Y - a.Y)
			if bAntinodeX >= 0 && bAntinodeX < m && bAntinodeY >= 0 && bAntinodeY < n {
				point := NewPoint(bAntinodeX, bAntinodeY)
				antinodes[point] = true
			}
		}
	}
}

func partTwo() int {
	antinodes := make(map[Point]bool)
	frequencies, m, n := parseInput()
	for _, points := range frequencies {
		findAntiNodesContinuous(antinodes, points, m, n)
	}

	return len(antinodes)

}

func findAntiNodesContinuous(antinodes map[Point]bool, points []Point, m, n int) {
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			distX := math.Pow(float64(a.X-b.X), 2)
			distY := math.Pow(float64(a.Y-b.Y), 2)
			dist := math.Sqrt(distX + distY)

			if dist < 1 {
				continue
			}

			offSetAX := a.X - b.X
			offSetAY := a.Y - b.Y
			aAntinodeX := a.X
			aAntinodeY := a.Y
			for aAntinodeX >= 0 && aAntinodeX < m && aAntinodeY >= 0 && aAntinodeY < n {
				point := NewPoint(aAntinodeX, aAntinodeY)
				antinodes[point] = true
				aAntinodeX += offSetAX
				aAntinodeY += offSetAY
			}

			offSetBX := b.X - a.X
			offSetBY := b.Y - a.Y
			bAntinodeX := b.X
			bAntinodeY := b.Y
			for bAntinodeX >= 0 && bAntinodeX < m && bAntinodeY >= 0 && bAntinodeY < n {
				point := NewPoint(bAntinodeX, bAntinodeY)
				antinodes[point] = true
				bAntinodeX += offSetBX
				bAntinodeY += offSetBY
			}
		}
	}

}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
