package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func part1(root *Node) int { 
	return max_sum(root, 100000)
	// return 0
}

func part2(root *Node) int {
	print("Space on disk ",70000000-root.total_size,"\n")
	print("Space needed ",30000000-(70000000-root.total_size),"\n")
	return min_total(root,30000000-(70000000-root.total_size), root.total_size)
	// return 0
}

func min_total(root *Node, min_size int, lowest int) int {
	total := lowest
	if (len(root.children) > 0) {
		for _, child := range root.children {
			total = min_total(child, min_size, total)
		}
	}
	if (root.total_size >= min_size && root.total_size < total) {
		
		total = root.total_size
	}
	return total
}

func max_sum(root *Node, max_size int) int {
	total := 0
	if (root.total_size <= max_size) {
		total += root.total_size
	}
	if (len(root.children) > 0) {
		for _, child := range root.children {
			total += max_sum(child, max_size)
		}
	}
	return total
}

type Node struct {
	name		string
	local_size	int
	total_size	int
	children	[]*Node
	parent		*Node
}

func handle_cd(curr *Node, dir string) *Node {
	
	if (dir == "..") {
		return curr.parent
	}
	for i, c := range curr.children {
		if (c.name == dir) {
			return curr.children[i]
		}
	}
	new_dir := Node {
		name:		dir,
		local_size:	0,
		total_size: 0,
		parent:		curr,
	}
	curr.children = append(curr.children[0:], &new_dir)
	return &new_dir
}

func printTree(root *Node, level int) {
	print("- Dir: ",root.name," : ",root.total_size,"\n")
	if (len(root.children) > 0) {
		for _, child := range root.children {
			printTree(child, level+1)
		}
	}
}

func calculate_total(root *Node) int {
	total_size := root.local_size
	if (len(root.children) > 0) {
		for _, child := range root.children {
			child_total := calculate_total(child)
			child.total_size = child_total
			total_size += child_total
		}
	}
	return total_size
}

func main() {
	// f, err := os.Open("./test.txt")
	f, err := os.Open("./day7_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	// total1 := 0
	// total2 := 0
	
	root := Node{
		name: 			"/",
		local_size: 	0,
		total_size:		0,
		parent:			nil,
	}
	curr := &root
	local_total := 0
	ls_flag := false
	
	for scanner.Scan() {
		row := scanner.Text()
		// print("ROW: ",row,"\n")
		if (row[0] == '$') {
			if (ls_flag) {
				curr.local_size = local_total
				local_total = 0
				ls_flag = false
			}
			// print("Running command... ",row[2:4],"\n")
			if (row[2:4] == "cd") {
				dir := row[5:]
				if (dir == "/") {
					curr = &root
					continue
				}
				curr = handle_cd(curr, dir)
			} else if (row[2:4] == "ls") {
				ls_flag = true
			}
		} else if (ls_flag) {
			data := strings.Split(row," ")
			count, err := strconv.Atoi(data[0])
			if err != nil {
				// handle error
			}
			if (count == 0) {
				
			}
			local_total += count
		} else {
			print("We shouldn't be here...\n")
			return
		}
	}
	if (ls_flag) {
		curr.local_size = local_total
		local_total = 0
		ls_flag = false
	}
	root.total_size = calculate_total(&root)
	
	
	fmt.Printf("Part1:\n%d\n",part1(&root))
	fmt.Printf("Part2:\n%d\n",part2(&root))
}
