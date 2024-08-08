package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	Value, X, Y, Depth int
	Children           []*Node
}

func createNode(x int, y int, d int, grid [][]int) Node {
	return Node{X: x, Y: y, Depth: d, Value: grid[x][y]}
}

type Queue []Node

func (q *Queue) Enqueue(n Node) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() (*Node, bool) {
	if len(*q) == 0 {
		return nil, false
	} else {
		n := (*q)[0]
		*q = (*q)[1:]
		return &n, true
	}
}

func isValid(x int, y int, m int, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func able(x int, y int, dX int, dY int, grid [][]int) bool {
	return grid[x][y] >= grid[dX][dY]-1
}

func able2(x int, y int, dX int, dY int, grid [][]int) bool {
	return grid[x][y]-1 <= grid[dX][dY]
}

// func dfs(x int, y int, grid [][]int, visitGrid [][]int, steps int, width int, height int) {
// 	X := [...]int{0, 1, 0, -1}
// 	Y := [...]int{-1, 0, 1, 0}
// 	visitGrid[x][y] = steps

// 	for i := range 4 {
// 		dX := x + X[i]
// 		dY := y + Y[i]

// 		if isValid(dX, dY, height, width) && willStep(x, y, dX, dY, grid, visitGrid, steps) {
// 			dfs(dX, dY, grid, visitGrid, steps+1, width, height)
// 		}
// 	}
// }

func bfs(root *Node, grid [][]int, visitGrid [][]bool, endX int, endY int, width int, height int) *Node {
	X := [...]int{0, 1, 0, -1}
	Y := [...]int{-1, 0, 1, 0}
	visitGrid[root.X][root.Y] = true
	queue := Queue{*root}

	for len(queue) != 0 {
		current, notEmpty := queue.Dequeue()
		if notEmpty {
			for i := range 4 {
				dX := current.X + X[i]
				dY := current.Y + Y[i]

				if isValid(dX, dY, height, width) && !visitGrid[dX][dY] && able(current.X, current.Y, dX, dY, grid) {
					child := createNode(dX, dY, current.Depth+1, grid)
					current.Children = append(current.Children, &child)
					visitGrid[dX][dY] = true
					queue.Enqueue(child)
					if dX == endX && dY == endY {
						return &child
					}
				}
			}
		}
	}
	return root
}

func bfs2(root *Node, grid [][]int, visitGrid [][]bool, val int, width int, height int) *Node {
	X := [...]int{0, 1, 0, -1}
	Y := [...]int{-1, 0, 1, 0}
	visitGrid[root.X][root.Y] = true
	queue := Queue{*root}

	for len(queue) != 0 {
		current, notEmpty := queue.Dequeue()
		if notEmpty {
			for i := range 4 {
				dX := current.X + X[i]
				dY := current.Y + Y[i]

				if isValid(dX, dY, height, width) && !visitGrid[dX][dY] && able2(current.X, current.Y, dX, dY, grid) {
					child := createNode(dX, dY, current.Depth+1, grid)
					current.Children = append(current.Children, &child)
					visitGrid[dX][dY] = true
					queue.Enqueue(child)
					if child.Value == val {
						return &child
					}
				}
			}
		}
	}
	return root
}

func Solve1() {
	f, err := os.Open("./12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var grid [][]int
	var endX, endY, count int
	var start Node

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		elevationRow := make([]int, len(row))
		for i, c := range row {
			if c == 'S' {
				start.X = count
				start.Y = i
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
	visitGrid := make([][]bool, height)
	for i := range height {
		visitGrid[i] = make([]bool, width)
	}
	// bfs
	n := bfs(&start, grid, visitGrid, endX, endY, width, height)
	answer = (*n).Depth

	fmt.Println("Day12 Solution 1: ", answer)
}

func Solve2() {
	f, err := os.Open("./12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var grid [][]int
	var count int
	var end Node

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		elevationRow := make([]int, len(row))
		for i, c := range row {
			if c == 'S' {
				c = 'a'
			} else if c == 'E' {
				c = 'z'
				end.X = count
				end.Y = i
			}
			elevationRow[i] = int(c)
		}
		grid = append(grid, elevationRow)
		count += 1
	}
	height := len(grid)
	width := len(grid[0])
	visitGrid := make([][]bool, height)
	for i := range height {
		visitGrid[i] = make([]bool, width)
	}
	// bfs
	n := bfs2(&end, grid, visitGrid, 97, width, height)
	answer = (*n).Depth

	fmt.Println("Day12 Solution 2: ", answer)
}
