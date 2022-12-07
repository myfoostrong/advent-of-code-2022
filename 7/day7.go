package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	// "strconv"
)

func part1() int { 
	
	return 0
}

func part2() int {
	
	return 0
}

func main() {
	// f, err := os.Open("./test.txt")
	f, err := os.Open("./day6_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	total1 := 0
	total2 := 0
	curr := ""
	char_count := 14 //4 for Part 1
	
	for scanner.Scan() {
		row := scanner.Text()
		
	}
	
	
	fmt.Printf("Part1:\n%d\n",total1)
	fmt.Printf("Part2:\n%d\n",total2)
}
