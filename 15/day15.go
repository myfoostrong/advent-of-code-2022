package day15

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

func parse(row string) []int {
	re := regexp.MustCompile("[-]?[0-9]+")
	coords := re.FindAllString(row, -1)
	var out []int
	for _, c := range coords {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, n)
	}
	return out
}

func getRange(sensor []int) int {
	return abs(sensor[0]-sensor[2]) + abs(sensor[1]-sensor[3])
}

func getMaxMin(sensor []int, minX, maxX, minY, maxY int) (int, int, int, int) {
	return min(min(sensor[0], minX), sensor[2]),
		max(max(sensor[0], maxX), sensor[2]),
		min(min(sensor[1], minY), sensor[3]),
		max(max(sensor[1], maxY), sensor[3])
}

func buildGrid(minX, maxX, minY, maxY int) [][]bool {
	y := (maxY - minY) + 1
	x := (maxX - minX) + 1
	grid := make([][]bool, y)
	for i := range y {
		grid[i] = make([]bool, x)
	}
	return grid
}

func normalize(sensorList [][]int, minX, minY int) [][]int {
	for i := range sensorList {
		sensorList[i][0] -= minX
		sensorList[i][1] -= minY
		sensorList[i][2] -= minX
		sensorList[i][3] -= minY
	}
	return sensorList
}

func markGrid(grid [][]bool, sensor []int, maxX, maxY int) [][]bool {
	dist := getRange(sensor)
	for dx := range dist + 1 {
		dy := dist - dx
		for i := range dy + 1 {
			if sensor[1]+i <= maxY {
				if sensor[0]+dx <= maxX {
					grid[sensor[1]+i][sensor[0]+dx] = true
				}
				if sensor[0]-dx >= 0 {
					grid[sensor[1]+i][sensor[0]-dx] = true
				}
			}
			if sensor[1]-i >= 0 {
				if sensor[0]+dx <= maxX {
					grid[sensor[1]-i][sensor[0]+dx] = true
				}
				if sensor[0]-dx >= 0 {
					grid[sensor[1]-i][sensor[0]-dx] = true
				}
			}
			printGrid(grid)
		}
		// grid[dy][dx] = true
	}
	return grid
}

func printGrid(grid [][]bool) {
	fmt.Println("\n")
	// for i := yMin; i < yMax; i++ {
	for y := range grid {
		for x := range grid[y] {
			a := "."
			if grid[y][x] {
				a = "#"
			}
			fmt.Print(a)
		}
		fmt.Print("\n")
	}
}

func Solve1() {
	f, err := os.Open("./15/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var minX, maxX, minY, maxY int
	var sensorList [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		sensor := parse(row)
		sensorList = append(sensorList, sensor)
		minX, maxX, minY, maxY = getMaxMin(sensor, minX, maxX, minY, maxY)
	}
	grid := buildGrid(minX, maxX, minY, maxY)
	sensorList = normalize(sensorList, minX, minY)
	maxX -= minX
	maxY -= minY
	for _, sensor := range sensorList {
		// printGrid(grid)
		grid = markGrid(grid, sensor, maxX, maxY)
	}

	for _, c := range grid[10] {
		if c {
			answer += 1
		}
	}
	fmt.Println("Day15 Solution 1: ", answer)
}

func Solve2() {

}
