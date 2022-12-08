package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func part1() int { 
	return 0
}

func part2() int {
	return 0
}

func main() {
	f, err := os.Open("./test.txt")
	// f, err := os.Open("./input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	// total1 := 0
	// total2 := 0
	
	for scanner.Scan() {
		row := scanner.Text()
		print("ROW: ",row,"\n")
	}
	
	
	fmt.Printf("Part1:\n%d\n",part1())
	fmt.Printf("Part2:\n%d\n",part2())
}
