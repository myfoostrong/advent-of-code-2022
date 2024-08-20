package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func build_stack(stack_msg string, stack_length int) []string {
	stack := make([]string, stack_length)

	return stack
}

func part1(stack []string, src int, dst int, count int) []string { 
	for i:= 0; i < count; i++ {
		stack[dst] = stack[src][:1] + stack[dst]
		stack[src] = stack[src][1:]
	}
	return stack
}

func part2(stack []string, src int, dst int, count int) []string {
	stack[dst] = stack[src][:count] + stack[dst]
	stack[src] = stack[src][count:]
	return stack
}

func main() {
	// f, err := os.Open("./test.txt")
	f, err := os.Open("./day5_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	part1_msg := ""
	part2_msg := ""
	counter := 0
	stack_height := 0
	stack_length := 0
	stack := []string{}
	stack1 := []string{}
	stack2 := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		if (counter == 0) {
			stack_length = (len(row) + 1)/4
			stack = make([]string, stack_length)
			stack1 = make([]string, stack_length)
			stack2 = make([]string, stack_length)
		}
		if (stack_height == 0) {				// This means we're still reading in the stack
			if (row[:4] == " 1  ") {			// We're at the bottom of the stack
				stack_height = counter + 1
			} else {
				for i := 0; i < stack_length; i++ {
					stack_column := row[i*4:i*4+3]
					stack_column = strings.ReplaceAll(stack_column," ","")
					stack_column = strings.ReplaceAll(stack_column,"[","")
					stack_column = strings.ReplaceAll(stack_column,"]","")
					stack[i] += stack_column					
				}
			}
		} else {
			if (len(row) > 0) {
				instr := strings.Split(row," ")
				count, err := strconv.Atoi(instr[1])
				src, err := strconv.Atoi(instr[3])
				dst, err := strconv.Atoi(instr[5])
				if err != nil {
					// handle error
				}
				src += -1
				dst += -1

				// Part 1
				stack1 = part1(stack1, src, dst, count)
				// for i, v := range stack1 {
				// 	fmt.Println(i,v)
				// }	

				// Part 2
				stack2 = part2(stack2, src, dst, count)	
				// for i, v := range stack2 {
				// 	fmt.Println(i,v)
				// }
				
			} else {
				copy(stack1,stack)
				copy(stack2,stack)
			}
		}
		counter += 1
	}
	for _, v := range stack1 {
		part1_msg += v[0:1]
	}
	for _, v := range stack2 {
		part2_msg += v[0:1]
	}
	
	fmt.Printf("Part1:\n%s\n",part1_msg)
	fmt.Printf("Part2:\n%s\n",part2_msg)
}
