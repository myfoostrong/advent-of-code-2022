package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	Value, X, Y int
	Children    []*Node
}

func isValid(x int, y int, m int, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func willStep(x int, y int, dX int, dY int, grid [][]int, visitGrid [][]int, steps int) bool {
	return visitGrid[dX][dY] > steps && grid[x][y] >= grid[dX][dY]-1
}

func dfs(x int, y int, grid [][]int, visitGrid [][]int, steps int, width int, height int) {
	X := [...]int{0, 1, 0, -1}
	Y := [...]int{-1, 0, 1, 0}
	visitGrid[x][y] = steps

	for i := range 4 {
		dX := x + X[i]
		dY := y + Y[i]

		if isValid(dX, dY, height, width) && willStep(x, y, dX, dY, grid, visitGrid, steps) {
			dfs(dX, dY, grid, visitGrid, steps+1, width, height)
		}
	}
}

func Solve1() {
	f, err := os.Open("./12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var grid [][]int
	var startX, startY, endX, endY, count int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		elevationRow := make([]int, len(row))
		for i, c := range row {
			if c == 'S' {
				startX = count
				startY = i
				c = 'a'
			} else if c == 'E' {
				c = 'z'
				endX = count
				endY = i
			}
			elevationRow[i] = int(c)
		}
		grid = append(grid, elevationRow)
		count += 1
	}
	height := len(grid)
	width := len(grid[0])
	visitGrid := make([][]int, height)
	for i := range height {
		visitGrid[i] = make([]int, width)
		for j := range width {
			// Initialize with more steps than are possible in given grid
			visitGrid[i][j] = width*height + 1
		}
	}
	dfs(startX, startY, grid, visitGrid, 0, width, height)
	answer = visitGrid[endX][endY]

	fmt.Println("Day12 Solution 1: ", answer)
}

func Solve2() {

}
