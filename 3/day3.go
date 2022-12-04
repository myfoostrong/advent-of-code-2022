package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_priority(c rune) int { 				// O(1)*5 == O(1)
	prio := int(c) - int('A')					// ASCII encoding: 'A': 65, 'a' : 97
	if(prio > 26) {								// Lower case, 1-26
		prio = prio % (int('a') - int('A'))		// i.e. % 32
	} else {									// Upper case 27-52
		prio = prio + 26
	}
	return prio
}

func part1(row string) int { 					// O(n) time, O(n+m) memory, where m is size of alphabet
	char_map := [52]int{}
	num_items := len(row)
	half_items := num_items / 2
	prio := 0
	for _, v := range row[:half_items] { 		// Catalog all items in first compartment
		prio = get_priority(v)
		char_map[prio] += 1
	}
	for _, v := range row[half_items:] { 		// Only check if the item existed in the first
		prio = get_priority(v)
		if(char_map[prio] > 0) {
			break								// Specifies one fail per pack, so we can assume we're done
		}
			
	}
	return prio	+ 1
}

func part2(group [3]string) int {
	char_map := [52]int{}
	badge := 0
	
	for i, row := range group[:2] { 
		for _, v := range row {
			prio := get_priority(v)
			if(char_map[prio] == i) {		// Always True for first row, only true in subsequent for first time char appears in row, if it existed in previous rows
				char_map[prio] = i+1		// Elves can carry duplicate badge items, so only want to increment once per row
			}			
		}
	}
	for _, v := range group[2] {
		prio := get_priority(v)
		if(char_map[prio] == 2) {			// We only care about chars in the first two lists
			badge = prio + 1
			break							// Save some cycles we're done
		}
	}
	return badge
}

func main() {
	f, err := os.Open("./day3_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	part1_total := 0
	part2_total := 0
	counter := 0
	group := [3]string{}
	for scanner.Scan() {
		row := scanner.Text()
		part1_total += part1(row)

		// Part 2
		group[counter] = row
		if (counter == 2) {
			part2_total += part2(group)
			counter = 0
		} else {
			counter += 1
		}
	}
	fmt.Printf("Part1 Total Priority %d \n",part1_total)
	fmt.Printf("Part2 Total Priority %d \n",part2_total)
}
