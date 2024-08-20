package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./day1_input.txt")
	if err != nil {
		// handle error
	}
	scanner := bufio.NewScanner(f)
	first_cal := 0
	first_elf := 0
	sec_elf := 0
	sec_cal := 0
	third_elf := 0
	third_cal := 0
	curr_cal := 0
	curr_elf := 0


	for scanner.Scan() {
		row := scanner.Text()
		if(row == "") {
			if(third_cal < curr_cal){
				if(sec_cal < curr_cal){
					if(first_cal < curr_cal){
						third_cal = sec_cal
						third_elf = sec_elf
						sec_cal = first_cal
						sec_elf = first_elf
						first_cal = curr_cal
						first_elf = curr_elf
						fmt.Printf("new 1st elf %d Cals: %d\n",curr_elf, curr_cal)
					} else {
						third_cal = sec_cal
						third_elf = sec_elf
						sec_cal = curr_cal
						sec_elf = curr_elf
						fmt.Printf("new 2nd elf %d Cals: %d\n",curr_elf, curr_cal)
					}
				} else {
					third_cal = curr_cal
					third_elf = curr_elf
					fmt.Printf("new 3rd elf %d Cals: %d\n",curr_elf, curr_cal)
				}
			}
			curr_elf++
			curr_cal = 0
		} else {
			cals, err := strconv.Atoi(row)
			if err != nil {
				// Add code here to handle the error!
			 }
			curr_cal += cals
		}
	}
	fmt.Printf("1st Elf: %d Cals: %d\n",first_elf+1, first_cal)
	fmt.Printf("2nd Elf: %d Cals: %d\n",sec_elf+1, sec_cal)
	fmt.Printf("3rd Elf: %d Cals: %d\n",third_elf+1, third_cal)
	fmt.Printf("Total Cals: %d\n",first_cal + sec_cal + third_cal)
}
