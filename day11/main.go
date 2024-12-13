package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() int {
	stones := parseInput()
	for range 25 {
		stones = blink(stones)
	}
	return len(stones)
}

func parseInput() []string {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		break
	}

	return strings.Split(line, " ")
}

func blink(stones []string) []string {
	newStones := []string{}
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			left := stone[:len(stone)/2]
			right := stone[len(stone)/2:]
			leftNum, err := strconv.Atoi(left)
			if err != nil {
				log.Fatal(err)
			}
			rightNum, err := strconv.Atoi(right)
			if err != nil {
				log.Fatal(err)
			}

			newStones = append(newStones, strconv.Itoa(leftNum))
			newStones = append(newStones, strconv.Itoa(rightNum))
		} else {
			val, err := strconv.Atoi(stone)
			if err != nil {
				log.Fatal(err)
			}
			newStones = append(newStones, strconv.Itoa(val*2024))
		}
	}

	return newStones
}

func partTwo() int {
	res := 0
	stones := parseInput()
	results := make(map[Memo]int)
	for _, stone := range stones {
		res += count(results, stone, 75)
	}
	return res
}

type Memo struct {
	Stone string
	Blink int
}

func NewMemo(stone string, blink int) Memo {
	return Memo{stone, blink}
}

func count(results map[Memo]int, stone string, remaining int) int {
	curr := NewMemo(stone, remaining)
	if remaining == 0 {
		results[curr] = 1
	}
	if freq, exists := results[curr]; exists {
		return freq
	}

	freq := 0
	if stone == "0" {
		freq += count(results, "1", remaining-1)
	} else if len(stone)%2 == 0 {
		left := stone[:len(stone)/2]
		right := stone[len(stone)/2:]
		leftNum, err := strconv.Atoi(left)
		if err != nil {
			log.Fatal(err)
		}
		rightNum, err := strconv.Atoi(right)
		if err != nil {
			log.Fatal(err)
		}
		freq += count(results, strconv.Itoa(leftNum), remaining-1)
		freq += count(results, strconv.Itoa(rightNum), remaining-1)
	} else {
		val, err := strconv.Atoi(stone)
		if err != nil {
			log.Fatal(err)
		}
		freq += count(results, strconv.Itoa(val*2024), remaining-1)
	}

	results[curr] = freq
	return results[curr]
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
