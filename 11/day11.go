package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y int
}

// Hello returns a greeting for the named person.
func Solve1() (int, error) {
	f, err := os.Open("./9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var tailVertices = make(map[Vertex]bool)
	var head, tail Vertex

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		move := strings.Split(row, " ")
		count, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatal(err)
		}
		for range count {
			head = moveHead(head, move[0])
			tail = followTail(head, tail, move[0])
			tailVertices[tail] = true
		}
	}

	// Return a greeting that embeds the name in a message.
	count := len(tailVertices)
	fmt.Println("Day9 Solution 1: ", count)

	return count, nil
}

func Solve2() (int, error) {
	f, err := os.Open("./9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var tailVertices = make(map[Vertex]bool)
	var rope [10]Vertex

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		move := strings.Split(row, " ")
		count, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatal(err)
		}
		for range count {
			for i := range len(rope) - 1 {
				if i == 0 {
					rope[i] = moveHead(rope[i], move[0])
				}
				rope[i+1] = followTail(rope[i], rope[i+1], move[0])
			}
			tailVertices[rope[9]] = true
		}
	}

	// Return a greeting that embeds the name in a message.
	count := len(tailVertices)
	fmt.Println("Day9 Solution 2: ", count)
	return count, nil
}

func absDiff(x int, y int) int {
	z := x - y
	if z < 0 {
		z = -z
	}
	return z
}

func abs(z int) int {
	if z < 0 {
		z = -z
	}
	return z
}

func moveHead(head Vertex, direction string) Vertex {
	switch direction {
	case "U":
		head.Y += 1
	case "D":
		head.Y -= 1
	case "L":
		head.X -= 1
	case "R":
		head.X += 1
	}
	return head
}

func followTail(head Vertex, tail Vertex, direction string) Vertex {
	// if absDiff(head.Y, tail.Y) > 1 {
	// 	tail.Y += (head.Y - tail.Y) / 2
	// }
	// if absDiff(head.X, tail.X) > 1 {
	// 	tail.X += (head.X - tail.X) / 2
	// }
	diffX := (head.X - tail.X)
	diffY := (head.Y - tail.Y)
	var stepX, stepY int

	if abs(diffX) == 2 {
		stepX = 1
		if diffY != 0 {
			stepY = 1
		}
	} else if abs(diffY) == 2 {
		stepY = 1
		if diffX != 0 {
			stepX = 1
		}
	}
	if diffX < 0 {
		stepX *= -1
	}
	if diffY < 0 {
		stepY *= -1
	}

	tail.X += stepX
	tail.Y += stepY

	// switch direction {
	// case "U":
	// 	if absDiff(head.Y, tail.Y) > 1 {
	// 		tail.Y += 1
	// 		tail.X += (head.X - tail.X)
	// 	}
	// case "D":
	// 	if absDiff(head.Y, tail.Y) > 1 {
	// 		tail.Y -= 1
	// 		tail.X += (head.X - tail.X)
	// 	}
	// case "L":
	// 	if absDiff(head.X, tail.X) > 1 {
	// 		tail.X -= 1
	// 		tail.Y += (head.Y - tail.Y)
	// 	}
	// case "R":
	// 	if absDiff(head.X, tail.X) > 1 {
	// 		tail.X += 1
	// 		tail.Y += (head.Y - tail.Y)
	// 	}
	// }
	return tail
}
