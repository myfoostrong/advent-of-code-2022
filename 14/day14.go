package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	f, err := os.Open("./14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var xMax, yMax int
	var rocks [][][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		coords_arr := strings.Split(row, " -> ")
		var points [][]int
		for _, coord_str := range coords_arr {
			point := strings.Split(coord_str, ",")
			var coords []int
			for i, c := range point {
				val, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}
				coords = append(coords, val)
				if i == 0 {
					if val > xMax {
						xMax = val
					}
				} else if i == 1 {
					if val > yMax {
						yMax = val
					}
				}
			}
			points = append(points, coords)
		}
		rocks = append(rocks, points)
	}
	answer = xMax + yMax
	fmt.Println("Day13 Solution 1: ", answer)
}

func Solve2() {

}
