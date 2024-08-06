package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve1() {
	f, err := os.Open("./12/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var grid []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	height := len(grid)
	width := len(grid[0])

	fmt.Println("Day12 Solution 1: ", answer)
}

func Solve2() {

}
