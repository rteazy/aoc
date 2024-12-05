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
	total := 0
	rules, orderings := parse()
	graph := buildGraph(rules)

	for _, order := range orderings {
		if valid := isValid(order, graph); valid {
			total += order[len(order)/2]
		}
	}

	return total
}

func buildGraph(rules []string) map[int][]int {
	graph := make(map[int][]int)
	for _, rule := range rules {
		vals := strings.Split(rule, "|")
		left, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatal(err)
		}

		graph[left] = append(graph[left], right)

	}
	return graph
}

func isValid(order []int, graph map[int][]int) bool {
	curr := order[0] // assume no index out of bounds
	i := 1
	for i < len(order) {
		found := false
		for _, node := range graph[curr] {
			if node == order[i] {
				found = true
				curr = node
			}
		}
		if !found {
			return false
		}
		i += 1
	}
	return true
}

func partTwo() int {
	total := 0
	rules, orderings := parse()
	graph := buildGraph(rules)

	for _, order := range orderings {
		if valid := isValid(order, graph); !valid {
			total += findOrder(order, graph)
		}
	}

	return total
}

func findOrder(order []int, graph map[int][]int) int {
	orderMap := make(map[int]bool)
	for _, val := range order {
		orderMap[val] = true
	}

	findRoot := func() (int, map[int]int) {
		indegrees := make(map[int]int)
		for _, val := range order {
			indegrees[val] = 0
		}

		for _, val := range order {
			for _, child := range graph[val] {
				if _, exists := orderMap[child]; exists {
					indegrees[child] += 1
				}
			}
		}

		for k, v := range indegrees {
			if v == 0 {
				return k, indegrees
			}
		}
		return -1, nil
	}

	root, indegrees := findRoot()
	queue := []int{root}
	save := []int{}
	for len(queue) > 0 {
		curr := queue[0]
		save = append(save, curr)
		queue = queue[1:]
		for _, neighbor := range graph[curr] {
			if _, exists := orderMap[neighbor]; exists {
				indegrees[neighbor] -= 1
				if indegrees[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return save[len(save)/2]
}

func parse() ([]string, [][]int) {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rules = append(rules, line)
	}

	orderings := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		ordering := []int{}
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			ordering = append(ordering, num)
		}
		orderings = append(orderings, ordering)
	}

	return rules, orderings
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
