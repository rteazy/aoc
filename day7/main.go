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
	res := 0
	valid := make(map[int]bool)

	var visit func(vals []int, index, total, target int)
	visit = func(vals []int, index, total, target int) {
		if index == len(vals)-1 {
			if total == target {
				valid[total] = true
			}
			return
		}
		nextIndex := index + 1
		nextVal := vals[nextIndex]
		add := total + nextVal
		multiply := total * nextVal
		if add <= target {
			visit(vals, nextIndex, add, target)
		}
		if multiply <= target {
			visit(vals, nextIndex, multiply, target)
		}
	}

	equations := parseInput()
	for _, equation := range equations {
		target, nums := equation.Target, equation.Nums
		visit(nums, 0, nums[0], target)
	}

	for num, _ := range valid {
		res += num
	}

	return res

}

type Equation struct {
	Target int
	Nums   []int
}

func NewEquation(target int, nums []int) Equation {
	return Equation{target, nums}
}

func parseInput() []Equation {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	equations := []Equation{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		targetStr, arrString := line[0], line[1]
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			log.Fatal(err)
		}
		valsString := strings.Split(arrString, " ")
		nums := []int{}
		for _, c := range valsString {
			val, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, val)
		}
		equation := NewEquation(target, nums)
		equations = append(equations, equation)
	}
	return equations
}

func partTwo() int {
	res := 0
	valid := make(map[int]bool)

	var visit func(vals []int, index, total, target int)
	visit = func(vals []int, index, total, target int) {
		if index == len(vals)-1 {
			if total == target {
				valid[total] = true
			}
			return
		}
		nextIndex := index + 1
		nextVal := vals[nextIndex]
		totalStr := strconv.Itoa(total)
		nextValStr := strconv.Itoa(nextVal)
		concatStr := totalStr + nextValStr
		concatVal, err := strconv.Atoi(concatStr)
		if err != nil {
			log.Fatal(err)
		}

		add := total + nextVal
		multiply := total * nextVal
		concat := concatVal

		if add <= target {
			visit(vals, nextIndex, add, target)
		}
		if multiply <= target {
			visit(vals, nextIndex, multiply, target)
		}
		if concat <= target {
			visit(vals, nextIndex, concat, target)
		}
	}

	equations := parseInput()
	for _, equation := range equations {
		target, nums := equation.Target, equation.Nums
		visit(nums, 0, nums[0], target)
	}

	for num, _ := range valid {
		res += num
	}

	return res
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
