package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"log"
	"sort"
)

func isError(err error) bool {
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
	return (err != nil)
}


func main() {
	var ip_file, error = os.Open("input-file")
	var read_calorie = 0
	var temp_elf_total int = 0
	var elf_food []int
	var top_three int = 0
	
	defer ip_file.Close()
	
	if isError(error) {
		return
	}

	// bufio allows us to read line by line
	scanner := bufio.NewScanner(ip_file)

	// as long as the line read is not EOF
	for scanner.Scan() {
		read_calorie, _ = strconv.Atoi(string(scanner.Text()))

		// if we don't come acrossa newline, total the elf's food
		if read_calorie != 0 {
		 	temp_elf_total = temp_elf_total + read_calorie

		} else {
			elf_food = append (elf_food, temp_elf_total)
			// uncomment for debug
			//fmt.Printf("total = %d\n", temp_elf_total) 
			temp_elf_total = 0
		}
	}

	// total the last elf's food separately
	elf_food = append (elf_food, temp_elf_total)

	// let's sort the food now
	sort.Slice(elf_food, func (p, q int) bool {
		return elf_food[p] > elf_food[q] })


	// The elf carrying the most food is...
	fmt.Printf ("Elf carrying highest food has = %d\n", elf_food[0])

	for i := 0; i < 3; i++ {
		top_three = top_three + elf_food[i]
	}
	
	fmt.Printf ("The top three elf food total = %d\n", top_three)
}
