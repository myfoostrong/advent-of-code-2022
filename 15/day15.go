package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parse(row string) ([]int, []int) {
	re := regexp.MustCompile("[0-9]+")
	coords := re.FindAllString(row, -1)
	var out []int
	for _, c := range coords {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, n)
	}
	return out[0:2], out[2:4]
}

func getRange(sensor, beacon []int) int {
	return abs(sensor[0]-beacon[0]) + abs(sensor[1]-beacon[1])
}

func getMaxMin(sensor []int, dist, minX, maxX, minY, maxY int) (int, int, int, int) {
	return min(sensor[0]-dist, minX),
		max(sensor[0]+dist, maxX),
		min(sensor[1]-dist, minY),
		max(sensor[1]+dist, maxY)
}

func buildGrid(minX, maxX, minY, maxY int) [][]bool {
	y := maxY - minY
	x := maxX - minX
	grid := make([][]bool, y)
	for i := range y {
		grid[i] = make([]bool, x)
	}
	return grid
}

func markGrid(grid [][]bool, sensor []int) [][]bool {
	dist := sensor[2]
	for dx := range dist {
		dy := dist - dx
		grid[dy][dx] = true
	}
	return grid
}

func Solve1() {
	f, err := os.Open("./14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var minX, maxX, minY, maxY int
	var sensorList [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		sensor, beacon := parse(row)
		dist := getRange(sensor, beacon)
		sensor = append(sensor, dist)
		sensorList = append(sensorList, sensor)
		minX, maxX, minY, maxY = getMaxMin(sensor, dist, minX, maxX, minY, maxY)
	}
	grid := buildGrid(minX, maxX, minY, maxY)
	for _, sensor := range sensorList {
		grid = markGrid(grid, sensor)
	}
	for _, c := range grid[10] {
		if c {
			answer += 1
		}
	}
	fmt.Println("Day14 Solution 1: ", answer)
}

func Solve2() {

}
