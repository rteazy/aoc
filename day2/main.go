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
	valid := 0
	reports := getReports()
	for _, report := range reports {
		if isSafe(report) {
			valid += 1
		}
	}

	return valid
}

func partTwo() int {
	valid := 0
	reports := getReports()
	for _, report := range reports {
		if isSafe(report) {
			valid += 1
		} else {
			// simulate level removal
			for i, _ := range report {
				clone := make([]int, len(report))
				copy(clone, report)
				newReport := append(clone[:i], clone[i+1:]...)
				if isSafe(newReport) {
					valid += 1
					break
				}
			}
		}
	}

	return valid
}

func getReports() [][]int {
	reports := [][]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strVals := strings.Split(line, " ")
		report := []int{}
		for _, c := range strVals {
			val, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, val)
		}

		reports = append(reports, report)
	}

	return reports
}

func isSafe(vals []int) bool {
	direction := 0
	n := len(vals)

	prev, curr := vals[0], vals[1]
	if prev == curr {
		return false
	} else if prev < curr {
		direction += 1
	} else {
		direction -= 1
	}

	for i := 1; i < n; i++ {
		curr := vals[i]
		dist := curr - prev
		if direction > 0 {
			if dist <= 0 || dist > 3 {
				return false
			}
		} else {
			if dist > -1 || dist < -3 {
				return false
			}
		}
		prev = curr
	}

	return true
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
