package day9

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
			tail = followTail(head, tail)
			tailVertices[tail] = true
		}
	}

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
				rope[i+1] = followTail(rope[i], rope[i+1])
			}
			tailVertices[rope[9]] = true
		}
	}

	// Return a greeting that embeds the name in a message.
	count := len(tailVertices)
	fmt.Println("Day9 Solution 2: ", count)
	return count, nil
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

func followTail(head Vertex, tail Vertex) Vertex {
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
	return tail
}
