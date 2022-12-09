package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	"strconv"
)

func part1() int { 
	return 0
}

func part2() int {
	return 0
}

func count_visible(row []string, half int) (int, int) {
	max := -1
	idx := 0
	count := 0
	for i, v := range row {
		height, _ := strconv.Atoi(v)
		if (height > max) {
			max = height
			idx = i
			count += 1
		} else if (i > half) {
			break
		}
	}
	return count, idx
}

func clear_visible(list [10]bool) [10]bool {
	for i := range list {
		list[i] = false
	}
	return list
}

func main() {
	f, err := os.Open("./test.txt")
	// f, err := os.Open("./input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)

	total1 := 0
	// total2 := 0
	// vertical_vis_right := [10][10]bool{}
	// vertical_vis_left := [10][10]bool{}
	// vertical_max := [10]int{}
	// vertical_last := [10]int{}
	count := 0
	trees := [][]bool{}
	grid := []string{}
	length := 0
	
	for scanner.Scan() {
		row := scanner.Text()
		print("ROW: ",row,"\n")

		if (count == 0) {
			length = len(row)
			trees = make([][]bool,length)
			for x := range trees {
				trees[x] = make([]bool, length)
			}
			grid = make([]string, length)
		}

		grid[count] = row

		max := -1
		// max_idx := 0
		// last := -1

		for i, v := range row {
			height, _ := strconv.Atoi(string(v))
			if (height > max) {
				// print("New max ",height,"\n")
				
				trees[count][i] = true
				// last = -1
				// for j:=(i-1); j>max_idx; j-- {					
				// 	// prev, _ := strconv.Atoi(string(row[j]))
				// 	// if(max == prev) {
				// 	// 	break
				// 	// }
				// 	print("Clearing: ",j,"\n")
				// 	trees[count][j] = false
				// }
				max = height
				// max_idx = i
			}
		}

		max = -1
		// for i:=(length-1);i>max_idx;i-- {
		// 	height, _ := strconv.Atoi(string(row[i]))
		// 	print("Height ", height,"\n")
		// 	if (height > max) {
		// 		trees[count][i] = true
		// 		max = height
		// 	}
		// }
			// else {
			// 	if (height > last) {
			// 		print("New forward max ",height,"\n")
			// 		for j:=(i-1); j>0; j-- {
			// 			prev, _ := strconv.Atoi(string(row[j]))
			// 			if (height > prev) {
			// 				print("Clearing: ",j,"\n")
			// 				trees[count][j] = false
			// 			} else {
			// 				break
			// 			}
			// 		}
			// 	} else if (height == last) {
			// 		continue
			// 	}					
			// 	// if(height == max) {
			// 	// 	continue
			// 	// }
			// 	trees[count][i] = true
			// 	last = height				
			// }			
			// }
		count += 1
	}

	// max := -1
	// max_idx := 0
	// last := -1

	// for i:=0;i<length;i++ {
	// 	max := -1
	// 	max_idx := 0
	// 	for j:=0;j<length;j++ {
	// 		height, _ := strconv.Atoi(string(grid[j][i]))
	// 		if (height > max) {
	// 			print("New max ",height,"\n")				
	// 			trees[j][i] = true
	// 			// last = -1
	// 			// for k:=(j-1); k>max_idx; k-- {
	// 			// 	print("Clearing: ",k,"\n")
	// 			// 	trees[k][i] = false
	// 			// }
	// 			max = height
	// 			max_idx = i
	// 		} 
	// 	}
	// 	max = -1
	// 	for j:=(length-1);j>max_idx;j-- {
	// 		height, _ := strconv.Atoi(string(grid[j][i]))
	// 		if (height > max) {
	// 			print("New max ",height,"\n")
	// 			trees[j][i] = true
	// 			max = height
	// 		}
	// 	}
	// }
			// else {
			// 	if (height > last) {
			// 		print("New forward max ",height,"\n")
			// 		for k:=(j-1); k>max_idx; k-- {
			// 			prev, _ := strconv.Atoi(string(grid[k][i]))
			// 			if (height > prev) {
			// 				print("Clearing: ",k,"\n")
			// 				trees[k][i] = false
			// 			} else {
			// 				break
			// 			}
			// 		}
			// 	} else if (height == last) {
			// 		continue
			// 	}
			// 	trees[j][i] = true
			// 	last = height				
			// }
			// }

	// print("\n\n")
	for i:=0;i<length;i++ {
		for j:=0;j<length;j++ {
			print(trees[i][j], "\t")
			if(i == 0 || j == 0 || i == (length-1) || j == (length-1)) {
				total1 +=1
			} else if(trees[i][j]) {
				total1 +=1
			}
		}
		print("\n")
	}
	
	
	fmt.Printf("Part1:\n%d\n",total1)
	fmt.Printf("Part2:\n%d\n",part2())
}
