package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve1() (int, error) {
	f, err := os.Open("./10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	x := 1
	cycle := 1
	dict := map[string]int{"noop": 1, "addx": 2}
	signal_cycles := []int{20, 60, 100, 140, 180, 220}
	var signals []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		instr := strings.Split(row, " ")

		cycle += dict[instr[0]]
		if cycle > signal_cycles[0] {
			signals = append(signals, (signal_cycles[0] * x))
			signal_cycles = signal_cycles[1:]
			if len(signal_cycles) == 0 {
				break
			}
		}
		if instr[0] == "addx" {
			val, err := strconv.Atoi(instr[1])
			if err != nil {
				log.Fatal(err)
			}
			x += val
		}

		if cycle == signal_cycles[0] {
			signals = append(signals, (signal_cycles[0] * x))
			signal_cycles = signal_cycles[1:]
			if len(signal_cycles) == 0 {
				break
			}
		}
	}

	// Return a greeting that embeds the name in a message.
	signal := 0
	for _, v := range signals {
		signal += v
	}
	fmt.Println("Day10 Solution 1: ", signal)

	return signal, nil
}
