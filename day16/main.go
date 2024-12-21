package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func partOne() int {
	grid, start := parseInput()
	return shortestPathToEnd(grid, start).distance
}

func shortestPathToEnd(grid [][]string, start Point) *Item {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	seen := make(map[Point]*Item)
	items := make(map[Point]*Item)

	startItem := &Item{
		point:    start,
		distance: 0,
	}
	items[start] = startItem
	heap.Push(&pq, startItem)
	prev := make(map[Point]map[Point]bool)
	end := []Point{}

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		delete(items, item.point)
		seen[item.point] = item
		dist := item.distance
		x, y, dx, dy := item.point.X, item.point.Y, item.point.dx, item.point.dy

		if grid[x][y] == "E" {
			end = append(end, item.point)
		}

		// counter-clockwise (-dj, di) ex.  (-1, 0)
		// i, j  => (di, dj) ex. (0, 1)
		// clockwise (dj, -di) => 1, 0
		neighbors := [][]int{
			{dist + 1, x + dx, y + dy, dx, dy},
			{dist + 1000, x, y, -dy, dx},
			{dist + 1000, x, y, dy, -dx},
		}
		for _, neighbor := range neighbors {
			alt, nX, nY, ndx, ndy := neighbor[0], neighbor[1], neighbor[2], neighbor[3], neighbor[4]
			if grid[nX][nY] == "#" {
				continue
			}
			neighborPoint := NewPoint(nX, nY, ndx, ndy)
			if _, visited := seen[neighborPoint]; visited {
				continue
			}
			if _, exists := items[neighborPoint]; !exists {
				neighborItem := &Item{point: neighborPoint, distance: math.MaxInt}
				items[neighborPoint] = neighborItem
			}
			if alt > items[neighborPoint].distance {
				continue
			}
			if alt < items[neighborPoint].distance {
				items[neighborPoint].distance = alt
				prev[neighborPoint] = make(map[Point]bool)
			}
			prev[neighborPoint][item.point] = true
			heap.Push(&pq, items[neighborPoint])
		}
	}

	best := end[0]
	queue := []Point{best}
	unique := make(map[Point]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		newNode := NewPoint(node.X, node.Y, 0, 0)
		unique[newNode] = true
		for pred := range prev[node] {
			queue = append(queue, pred)
		}
	}
	fmt.Println(len(unique))

	return seen[end[0]]
}

type Point struct {
	X, Y, dx, dy int
}

func NewPoint(x, y, dx, dy int) Point {
	return Point{x, y, dx, dy}
}

// Adapted from go docs PriorityQueue example: https://pkg.go.dev/container/heap#Interface

type Item struct {
	point    Point
	distance int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func parseInput() ([][]string, Point) {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	var start Point
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == "S" {
				start = NewPoint(i, j, 0, 1)
			}
		}
	}

	return grid, start
}

func partTwo() int {
	// grid, start := parseInput()
	// end := shortestPathToEnd(grid, start)
	// best := end[0]
	// queue := []Point{best}
	// unique := make(map[Point]bool)
	// for len(queue) > 0 {
	// 	node := queue[0]
	// 	queue = queue[1:]
	// 	newNode := NewPoint(node.X, node.Y, 0, 0)
	// 	unique[newNode] = true
	// 	for pred := range prev[node] {
	// 		queue = append(queue, pred)
	// 	}
	// }
	// fmt.Println(unique)

	// fmt.Printf("Count: %d\n", len(unique))
	return 0
}

func main() {
	fmt.Println(partOne())
}
