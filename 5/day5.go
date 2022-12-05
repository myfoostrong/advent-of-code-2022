package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(grid [4]int) int { 
	return 0
}

func part2(grid [4]int) int {
	return 0
}

func main() {
	// f, err := os.Open("./test.txt")
	f, err := os.Open("./day4_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	part1_total := 0
	part2_total := 0
	for scanner.Scan() {
		row := scanner.Text()
		grid := [4]int{}
		for i, zone := range strings.Split(row, ",") {				// Convert input to int array
			for j, coord := range strings.Split(zone, "-") {
				n, err := strconv.Atoi(coord)
				if err != nil {
					// handle error
				}
				grid[i*2 + j] = n
			}
		}
		part1_total += part1(grid)
		part2_total += part2(grid)
		
	}
	fmt.Printf("Part1 Total Priority %d \n",part1_total)
	fmt.Printf("Part2 Total Priority %d \n",part2_total)
}
