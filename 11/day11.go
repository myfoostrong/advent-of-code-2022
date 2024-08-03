package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Monkey struct {
	ID, TestNum, Pass, Fail int
	Items                   []int
	OperationStr            []string
}

func (m Monkey) test(x int) bool {
	return x%m.TestNum == 0
}

func (m Monkey) operation(item int) (int, error) {
	var x int
	if m.OperationStr[2] == "old" {
		x = item
	} else {
		val, err := strconv.Atoi(m.OperationStr[2])
		if err != nil {
			log.Fatal(err)
		}
		x = val
	}
	if m.OperationStr[1] == "*" {
		return item * x, nil
	} else if m.OperationStr[1] == "+" {
		return item + x, nil
	} else {
		log.Fatal("Invalid math operation")
	}
	log.Fatal("Bad monkey")
	return 0, nil
}

func newMonkey(input []string) Monkey {

}

func Solve1() {
	f, err := os.Open("./10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var monkeys []Monkey
	var input []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			monkeys = append(monkeys, newMonkey(input))
			input = nil
		} else {
			input = append(input, row)
		}
	}

	fmt.Println("Day10 Solution 1: ", answer)
}
