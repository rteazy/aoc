package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() int {
	res := 0
	blocks := parseInput()
	moveFreeSpaceToEnd(blocks)
	for i, block := range blocks {
		if block.IsFreeSpace {
			break
		}
		res += i * block.Id
	}
	return res
}

type Block struct {
	Id          int
	IsFreeSpace bool
}

func NewBlock(id int, isFreeSpace bool) Block {
	return Block{id, isFreeSpace}
}

func parseInput() []Block {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	blocks := []Block{}
	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}

	id := 0
	for i, c := range line {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		if i%2 == 0 {
			for range val {
				block := NewBlock(id, false)
				blocks = append(blocks, block)
			}
			id += 1
		} else {
			for range val {
				block := NewBlock(-1, true)
				blocks = append(blocks, block)
			}
		}

	}
	return blocks
}

func moveFreeSpaceToEnd(blocks []Block) {
	i, j := 0, len(blocks)-1
	for i < j {
		for i < j && !blocks[i].IsFreeSpace {
			i += 1
		}
		for i < j && blocks[j].IsFreeSpace {
			j -= 1
		}
		if i < j {
			blocks[i], blocks[j] = blocks[j], blocks[i]
		}
	}
}

func partTwo() int {
	res := 0
	blocks := parseInput()
	moveFile(blocks)
	for i, block := range blocks {
		if !block.IsFreeSpace {
			res += i * block.Id
		}
	}
	return res
}

func moveFile(blocks []Block) {
	j := len(blocks) - 1
	for j > 0 {
		if blocks[j].IsFreeSpace {
			j -= 1
			continue
		}
		currID := blocks[j].Id
		fileLength := 0
		for idx := j; idx > 0 && blocks[idx].Id == currID; idx-- {
			fileLength += 1
		}

		// search for file block on left
		for i := 0; i < j; i++ {
			if blocks[i].IsFreeSpace {
				freeSpace := 0
				for idx := i; idx < j && blocks[idx].IsFreeSpace; idx++ {
					freeSpace += 1
				}

				// perform swap if enough space
				if freeSpace >= fileLength {
					for k := range fileLength {
						blocks[i+k], blocks[j-k] = blocks[j-k], blocks[i+k]
					}
					break
				}
			}
		}

		j -= fileLength
	}
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
