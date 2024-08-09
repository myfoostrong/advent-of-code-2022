package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func parse(s string) []any {
	packet := []any{}
	json.Unmarshal([]byte(s), &packet)
	return packet
}

func getType(a any) (int, []any, bool) {
	switch val := a.(type) {
	default:
		log.Fatal("Uknown type, panic!")
		return 0, nil, false
	case float64:
		return int(val), nil, true
	case int:
		return val, nil, true
	case []any:
		return 0, val, false
	}
}

func test(left any, right any) int {
	leftInt, leftList, isLeftNum := getType(left)
	rightInt, rightList, isRightNum := getType(right)
	if !isLeftNum && !isRightNum {
		return testList(leftList, rightList)
	} else if !isLeftNum && isRightNum {
		return testList(leftList, []any{rightInt})
	} else if isLeftNum && !isRightNum {
		return testList([]any{leftInt}, rightList)
	} else if isLeftNum && isRightNum {
		return testNum(leftInt, rightInt)
	}
	panic("wrong")
}

func testList(left []any, right []any) int {
	var minLen int
	if len(left) < len(right) {
		minLen = len(left)
	} else {
		minLen = len(right)
	}
	for i := 0; i < minLen; i++ {
		if cmp := test(left[i], right[i]); cmp != 0 {
			return cmp
		}
	}
	return testNum(len(left), len(right))
}

func testNum(left int, right int) int {
	if left < right {
		return -1
	} else if left == right {
		return 0
	} else if left > right {
		return 1
	}
	panic("wrong")
}

func Solve1() {
	f, err := os.Open("./13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
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
		p1 := parse(pair[0])
		p2 := parse(pair[1])
		if test(p1, p2) == -1 {
			ordered = append(ordered, i+1)
		}
	}
	for _, x := range ordered {
		answer += x
	}
	fmt.Println("Day13 Solution 1: ", answer)
}

func quicksort(a []any, lo, hi int) []any {
	if lo >= 0 && hi >= 0 && lo < hi {
		var p int
		a, p = partition(a, lo, hi)
		a = quicksort(a, lo, p-1)
		a = quicksort(a, p+1, hi)
	}
	return a
}

func partition(a []any, lo, hi int) ([]any, int) {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if test(a[j], pivot) == -1 {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return a, i
}

func Solve2() {
	f, err := os.Open("./13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var answer int
	var packets []any

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		if row != "" {
			packets = append(packets, parse(row))
		}
	}
	packets = append(packets, parse("[[6]]"))
	packets = append(packets, parse("[[2]]"))

	packets = quicksort(packets, 0, len(packets)-1)
	var ordered []int
	for i, p := range packets {
		if _, y1, z := getType(p); !z && len(y1) == 1 {
			for _, a := range y1 {
				if _, y2, z := getType(a); !z && len(y2) == 1 {
					for _, a1 := range y2 {
						if x, _, isNum := getType(a1); isNum {
							if x == 2 || x == 6 {
								ordered = append(ordered, i+1)
							}
						}
					}
				}
			}
			// if _, y2, z := getType(y1); !z && len(y2) == 1 {
			// 	if _, y3, z := getType(y2); !z && len(y3) == 1 {
			// 		if a, _, c := getType(y3); c {
			// if a == 2 || a == 6 {
			// 	ordered = append(ordered, i)
			// }
			// 		}
			// 	}
			// }
		}
	}
	answer = ordered[0] * ordered[1]
	fmt.Println("Day13 Solution 2: ", answer)
}
