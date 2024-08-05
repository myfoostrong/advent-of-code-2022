package main

import (
	"aoc22/day10"
	"aoc22/day11"
	"aoc22/day9"
	"log"
)

func main() {
	log.SetFlags(0)

	day9.Solve1()
	day9.Solve2()

	day10.Solve1()
	day10.Solve2()

	day11.Solve1()
	day11.Solve2()
}
