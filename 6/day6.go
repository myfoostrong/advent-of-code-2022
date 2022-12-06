package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	// "strconv"
)

func build_stack(stack_msg string, stack_length int) []string {
	stack := make([]string, stack_length)

	return stack
}

func part1() int { 
	
	return 0
}

func part2() int {
	
	return 0
}

func main() {
	f, err := os.Open("./test.txt")
	// f, err := os.Open("./day6_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	part1_marker := 0
	part2_marker := ""
	curr := ""
	
	for scanner.Scan() {
		row := scanner.Text()
		signal := [26]int{}
		curr = row[:4]
		print(curr,"\n")
		for i, char := range row {
			if (i > 3) {
				c := int(char) - int('a')
				signal[c] += 1
				if signal[c] != 1 {
					
				} else {
					print(i,"\n")
					break
				}
			}
		}
	}
	
	fmt.Printf("Part1:\n%s\n",part1_marker)
	fmt.Printf("Part2:\n%s\n",part2_marker)
}
