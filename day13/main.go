package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func partOne() int {
	res := 0
	claws := parseInput()
	for _, claw := range claws {
		res += findMinTokens(claw)
	}
	return res
}

func parseInput() []Claw {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	claws := []Claw{}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("X\\+(\\d+), Y\\+(\\d+)")
	r2, _ := regexp.Compile("X=(\\d+), Y=(\\d+)")
	for scanner.Scan() {
		buttonALine := scanner.Text()
		buttonAStr := r.FindStringSubmatch(buttonALine)
		buttonAX, err := strconv.Atoi(buttonAStr[1])
		if err != nil {
			log.Fatal(err)
		}
		buttonAY, err := strconv.Atoi(buttonAStr[2])
		if err != nil {
			log.Fatal(err)
		}
		buttonA := NewPoint(buttonAX, buttonAY)

		scanner.Scan()
		buttonBLine := scanner.Text()
		buttonBStr := r.FindStringSubmatch(buttonBLine)
		buttonBX, err := strconv.Atoi(buttonBStr[1])
		if err != nil {
			log.Fatal(err)
		}
		buttonBY, err := strconv.Atoi(buttonBStr[2])
		if err != nil {
			log.Fatal(err)
		}
		buttonB := NewPoint(buttonBX, buttonBY)

		scanner.Scan()
		prizeLine := scanner.Text()
		prizeStr := r2.FindStringSubmatch(prizeLine)
		prizeX, err := strconv.Atoi(prizeStr[1])
		if err != nil {
			log.Fatal(err)
		}
		prizeY, err := strconv.Atoi(prizeStr[2])
		if err != nil {
			log.Fatal(err)
		}
		prize := NewPoint(prizeX, prizeY)

		claw := NewClaw(buttonA, buttonB, prize)
		claws = append(claws, claw)
		scanner.Scan()
	}

	return claws
}

func findMinTokens(claw Claw) int {
	minPrice := math.MaxInt64
	prizeFound := [][]bool{}
	for range 100 {
		prizeI := []bool{}
		for range 100 {
			prizeI = append(prizeI, false)
		}
		prizeFound = append(prizeFound, prizeI)
	}

	var findPrize func([][]bool, Claw)
	findPrize = func(found [][]bool, curr Claw) {
		for i := range len(found) {
			for j := range len(found[0]) {
				x := curr.ButtonA.X*i + curr.ButtonB.X*j
				y := curr.ButtonA.Y*i + curr.ButtonB.Y*j
				if x == claw.Prize.X && y == claw.Prize.Y {
					found[i][j] = true
				}
			}
		}
	}

	findPrize(prizeFound, claw)

	for i := range len(prizeFound) {
		for j := range len(prizeFound[0]) {
			if prizeFound[i][j] {
				minPrice = min(minPrice, 3*i+j)
			}
		}
	}

	if minPrice == math.MaxInt64 {
		return 0
	} else {
		return minPrice
	}

}

type Claw struct {
	ButtonA, ButtonB, Prize Point
}

func NewClaw(a, b, p Point) Claw {
	return Claw{a, b, p}
}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func partTwo() int {
	res := 0
	claws := parseInput()
	for _, claw := range claws {
		res += findMinTokensTwoEquations(claw)
	}
	return res
}

func findMinTokensTwoEquations(claw Claw) int {
	prizeX := claw.Prize.X + 10000000000000
	prizeY := claw.Prize.Y + 10000000000000

	// Solve system of two equations
	// claw.ButtonA.X * x + claw.ButtonB.X * y = prizeX
	// claw.ButtonA.Y * x + claw.ButtonB.Y * y = prizeY

	dx := claw.ButtonA.Y * claw.ButtonB.X
	px := claw.ButtonA.Y * prizeX
	dy := -claw.ButtonA.X * claw.ButtonB.Y
	py := -claw.ButtonA.X * prizeY

	y := dx + dy
	p := px + py
	if p%y != 0 {
		return 0
	}

	j := p / y
	res := prizeX - (claw.ButtonB.X * j)
	if res%claw.ButtonA.X != 0 {
		return 0
	}
	i := res / claw.ButtonA.X

	return 3*i + j
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
