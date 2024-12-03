package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne() int {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r, err := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
		if err != nil {
			log.Fatal(err)
		}

		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			left, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			total += left * right

		}
	}

	return total
}

func partTwo() int {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		r, err := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|do\\(\\)|don't\\(\\)")
		if err != nil {
			log.Fatal(err)
		}

		submatches := r.FindAllStringSubmatch(line, -1)
		for _, submatch := range submatches {
			expr := submatch[0]
			if strings.HasPrefix(expr, "do()") {
				enabled = true
			} else if strings.HasPrefix(expr, "don't()") {
				enabled = false
			} else if enabled {
				left, err := strconv.Atoi(submatch[1])
				if err != nil {
					log.Fatal(err)
				}
				right, err := strconv.Atoi(submatch[2])
				if err != nil {
					log.Fatal(err)
				}
				total += left * right
			}
		}
	}

	return total
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
