package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	f, err := os.Open("./14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var xMin, yMin, xMax, yMax int = 500, 500, 0, 0
	var rocks [][][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		coords_arr := strings.Split(row, " -> ")
		var points [][]int
		for _, coord_str := range coords_arr {
			point := strings.Split(coord_str, ",")
			var coords []int
			for i, c := range point {
				val, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}
				coords = append(coords, val)
				if i == 0 {
					if val > xMax {
						xMax = val
					}
					if val < xMin {
						xMin = val
					}
				} else if i == 1 {
					if val > yMax {
						yMax = val
					}
					if val < yMin {
						yMin = val
					}
				}
			}
			points = append(points, coords)
		}
		rocks = append(rocks, points)
	}
	xMax += 4
	xMin -= 1
	yMax += 2
	yMin -= 1
	grid := make([][]bool, yMax)
	for i := range yMax {
		grid[i] = make([]bool, xMax)
	}
	for _, row := range rocks {
		var last []int
		for i, point := range row {
			if i == 0 {
				last = point
				grid[point[1]][point[0]] = true
				continue
			}

			if last[0] == point[0] {
				dir := point[1] - last[1]
				if dir < 0 {
					for j := last[1]; j > point[1]-1; j-- {
						grid[j][point[0]] = true
					}
				} else if dir == 0 {
					log.Fatal("Shouldn't be zero")
				} else {
					for j := last[1]; j < point[1]+1; j++ {
						grid[j][point[0]] = true
					}
				}
			} else if last[1] == point[1] {
				dir := point[0] - last[0]
				if dir < 0 {
					for j := last[0]; j > point[0]-1; j-- {
						grid[point[1]][j] = true
					}
				} else if dir == 0 {
					log.Fatal("Shouldn't be zero")
				} else {
					for j := last[0]; j < point[0]+1; j++ {
						grid[point[1]][j] = true
					}
				}
			}
			last = point
		}
	}
	// printGrid(grid, xMin, xMax, 0, yMax)
	var count, drop int
sand:
	x := 500
	for y := range yMax {
		if grid[y][x] {
			if x > 0 && grid[y][x-1] {
				if grid[y][x+1] {
					grid[y-1][x] = true
					// printGrid(grid, xMin, xMax, 0, 55)
					count += 1
					goto sand
				} else {
					x += 1
				}
			} else {
				x -= 1
			}
		}
	}
	drop += 1
	if drop < 2 {
		goto sand
	}
	answer = count
	// printGrid(grid, xMin, xMax, 0, yMax)
	fmt.Println("Day14 Solution 1: ", answer)
}

func printGrid(grid [][]bool, xMin, xMax, yMin, yMax int) {
	fmt.Println("\n")
	for i := yMin; i < yMax; i++ {
		for j, c := range grid[i][xMin:xMax] {
			a := "."
			if c {
				a = "#"
			}
			if i == 0 && j == 500-xMin {
				a = "x"
			}
			fmt.Print(a)
		}
		fmt.Print("\n")
	}
}

func Solve2() {
	f, err := os.Open("./14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var xMin, yMin, xMax, yMax int = 500, 500, 0, 0
	var rocks [][][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		coords_arr := strings.Split(row, " -> ")
		var points [][]int
		for _, coord_str := range coords_arr {
			point := strings.Split(coord_str, ",")
			var coords []int
			for i, c := range point {
				val, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}
				coords = append(coords, val)
				if i == 0 {
					if val > xMax {
						xMax = val
					}
					if val < xMin {
						xMin = val
					}
				} else if i == 1 {
					if val > yMax {
						yMax = val
					}
					if val < yMin {
						yMin = val
					}
				}
			}
			points = append(points, coords)
		}
		rocks = append(rocks, points)
	}
	xMax += 10000
	xMin -= 1
	yMax += 4
	yMin -= 1
	grid := make([][]bool, yMax)
	for i := range yMax {
		grid[i] = make([]bool, xMax)
	}
	for _, row := range rocks {
		var last []int
		for i, point := range row {
			if i == 0 {
				last = point
				grid[point[1]][point[0]] = true
				continue
			}

			if last[0] == point[0] {
				dir := point[1] - last[1]
				if dir < 0 {
					for j := last[1]; j > point[1]-1; j-- {
						grid[j][point[0]] = true
					}
				} else if dir == 0 {
					log.Fatal("Shouldn't be zero")
				} else {
					for j := last[1]; j < point[1]+1; j++ {
						grid[j][point[0]] = true
					}
				}
			} else if last[1] == point[1] {
				dir := point[0] - last[0]
				if dir < 0 {
					for j := last[0]; j > point[0]-1; j-- {
						grid[point[1]][j] = true
					}
				} else if dir == 0 {
					log.Fatal("Shouldn't be zero")
				} else {
					for j := last[0]; j < point[0]+1; j++ {
						grid[point[1]][j] = true
					}
				}
			}
			last = point
		}
	}
	for i := range xMax {
		grid[yMax-2][i] = true
	}
	// printGrid(grid, xMin, xMax, 0, yMax)
	var count int
sand:
	x := 500
	for y := range yMax {
		if grid[y][x] {
			if x > 0 && x < xMax-1 && grid[y][x-1] {
				if grid[y][x+1] {
					grid[y-1][x] = true
					// printGrid(grid, xMin, xMax, 0, yMax)
					count += 1
					if y == 1 && x == 500 {
						break
					}
					goto sand
				} else {
					x += 1
				}
			} else {
				x -= 1
			}
		}
	}
	// printGrid(grid, xMin, xMax, 0, yMax)
	answer = count
	fmt.Println("Day14 Solution 2: ", answer)
}
