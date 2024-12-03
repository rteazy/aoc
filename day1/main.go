package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partOne(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	total := 0
	n := len(right)
	for i := 0; i < n; i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}
		total += dist
	}

	return total
}

func partTwo(left, right []int) int {
	res := 0
	slices.Sort(left)
	counts := make(map[int]int)
	for _, val := range right {
		if _, exists := counts[val]; !exists {
			counts[val] = 1
		} else {
			counts[val] += 1
		}

	}
	for _, val := range left {
		if _, exists := counts[val]; exists {
			res += val * counts[val]
		}
	}

	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left, right := []int{}, []int{}
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, "   ")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatalf("fail to parse")
		}
		left = append(left, a)

		b, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatalf("fail to parse")
		}
		right = append(right, b)
	}

	resPartOne := partOne(left, right)
	fmt.Println(resPartOne)

	resPartTwo := partTwo(left, right)
	fmt.Println(resPartTwo)
}
