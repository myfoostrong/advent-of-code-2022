package day11

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	ID, Pass, Fail int
	TestNum        float64
	Items          []float64
	Operation      []string
}

func (m Monkey) test(x float64) int {
	if math.Mod(x, m.TestNum) == 0 {
		return m.Pass
	}
	return m.Fail
}

func (m Monkey) operation(item float64) float64 {
	var x float64
	if m.Operation[2] == "old" {
		x = item
	} else {
		val, err := strconv.ParseFloat(m.Operation[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		x = val
	}
	if m.Operation[1] == "*" {
		return item * x
	} else if m.Operation[1] == "+" {
		return item + x
	} else {
		log.Fatal("Invalid math operation")
	}
	log.Fatal("Bad monkey")
	return 0
}

func newMonkey(input []string) Monkey {
	// Monkey ID
	idRow := strings.Split(input[0], " ")[1]
	id, err := strconv.Atoi(idRow[:len(idRow)-1])
	if err != nil {
		log.Fatal(err)
	}

	// Starting Items
	itemsRow := strings.Split(strings.Split(input[1], ": ")[1], ", ")
	var items []float64
	for _, itemStr := range itemsRow {
		item, err := strconv.ParseFloat(itemStr, 64)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	//Operation
	operation := strings.Split(strings.Split(input[2], "= ")[1], " ")

	//Test
	test, err := strconv.ParseFloat(strings.Split(input[3], "by ")[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	//Pass
	pass, err := strconv.Atoi(strings.Split(input[4], "monkey ")[1])
	if err != nil {
		log.Fatal(err)
	}

	//Fail
	fail, err := strconv.Atoi(strings.Split(input[5], "monkey ")[1])
	if err != nil {
		log.Fatal(err)
	}

	return Monkey{
		ID:        id,
		Items:     items,
		Operation: operation,
		TestNum:   test,
		Pass:      pass,
		Fail:      fail,
	}

}

func Solve1() {
	f, err := os.Open("./11/input.txt")
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
	monkeys = append(monkeys, newMonkey(input))

	inspections := make([]int, len(monkeys))
	for range 20 {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				newItem := math.Floor(monkey.operation(item) / 3)
				target := monkey.test(newItem)
				monkeys[target].Items = append(monkeys[target].Items, newItem)
				inspections[i] += 1
			}
			monkeys[i].Items = nil
		}
	}

	var first, second int
	for _, x := range inspections {
		if x > first {
			second = first
			first = x
			continue
		} else if x > second {
			second = x
		}
	}
	answer = first * second

	fmt.Println("Day11 Solution 1: ", answer)
}

func Solve2() {
	f, err := os.Open("./11/test.txt")
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
	monkeys = append(monkeys, newMonkey(input))

	inspections := make([]int, len(monkeys))
	for j := range 10000 {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				newItem := monkey.operation(item)
				target := monkey.test(newItem)
				monkeys[target].Items = append(monkeys[target].Items, newItem)
				inspections[i] += 1
			}
			monkeys[i].Items = nil
		}
		if j == 999 {
			answer = j
		}
	}

	var first, second int
	for _, x := range inspections {
		if x > first {
			second = first
			first = x
			continue
		} else if x > second {
			second = x
		}
	}
	answer = first * second

	fmt.Println("Day11 Solution 1: ", answer)
}
