package main

import (
	"aoc22/day11"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// day9.Solve1()
	// day9.Solve2()

	day11.Solve1()
	day11.Solve2()
	// If an error was returned, print it to the console and
	// exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If no error was returned, print the returned message
	// to the console.
	// message := fmt.Sprintf("Day1 ", message)
	// fmt.Println("Day9 Solution 1: ", message)
}
