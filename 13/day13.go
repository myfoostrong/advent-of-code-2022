package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getList(s string) ([]string, string) {
	ends := 1
	for i, c := range s {
		if c == '[' {
			ends += 1
			continue
		} else if c == ']' {
			ends -= 1
		}
		if ends == 0 {
			list := []string{s[:i]}
			return list, s[i:]
		}
	}
	return nil, s
}

func Solve1() {
	f, err := os.Open("./12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var count int
	var pair []string
	var packets [][]string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			packets = append(packets, pair)
			pair = nil
			continue
		}
		pair = append(pair, row)
	}
	var ordered []int
	for i, pair := range packets {
		p1 := pair[0]
		p2 := pair[1]
		for _, c1 := range p1[1 : len(p1)-1] {
			if c1 == '[' {
				getList(p1)
			}
		}
	}

	fmt.Println("Day12 Solution 1: ", answer)
}

func Solve2() {

}
