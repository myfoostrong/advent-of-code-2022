package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./day2_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	outcomes := map[string]map[string]int{}
	outcomes["A"] = make(map[string]int)
	outcomes["B"] = make(map[string]int)	
	outcomes["C"] = make(map[string]int)

	// outcomes["A"]["X"] = 4
	// outcomes["A"]["Y"] = 8
	// outcomes["A"]["Z"] = 3

	// outcomes["B"]["X"] = 1
	// outcomes["B"]["Y"] = 5
	// outcomes["B"]["Z"] = 9

	// outcomes["C"]["X"] = 7
	// outcomes["C"]["Y"] = 2
	// outcomes["C"]["Z"] = 6

	//part 2
	outcomes["A"]["X"] = 3
	outcomes["A"]["Y"] = 4
	outcomes["A"]["Z"] = 8

	outcomes["B"]["X"] = 1
	outcomes["B"]["Y"] = 5
	outcomes["B"]["Z"] = 9

	outcomes["C"]["X"] = 2
	outcomes["C"]["Y"] = 6
	outcomes["C"]["Z"] = 7

	total := 0
	for scanner.Scan() {
		row := scanner.Text()
		plays := strings.Fields(row)
		total += outcomes[plays[0]][plays[1]]
		
	}
	fmt.Printf("aa %d \n",total)
}
