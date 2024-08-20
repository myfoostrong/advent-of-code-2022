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

	marker := 0
	curr := ""
	char_count := 14 //4 for Part 1
	
	for scanner.Scan() {
		row := scanner.Text()
		print(curr,"\n")
		for i := range row {
			match := false
			if (i > char_count - 1) {
				signal := row[i-char_count:i]
				// print("Checking signal ",signal,"\n")
				for j,v := range signal {
					for k:=j+1; k<char_count; k++ {
						// fmt.Printf("Does %s match %s?\n",signal[k],v)
						if (string(signal[k]) == string(v)) {
							// print("Yes\n")
							match = true
							break
						}
					}
					if (match) {
						break
					}
				}
				if (!match) {
					marker = i
					break
				}
			}
		}
		// fmt.Printf("Part1:\n%d\n",marker)
		fmt.Printf("Part2:\n%d\n",marker)
	}
	
	
	
}
