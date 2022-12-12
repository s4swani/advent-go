package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"strconv"
)

func isError(err error) bool {

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return (err != nil)
}

func main() {
	var full_overlap_count int = 0
	var partial_overlap_count int = 0

	var ip_file, err = os.Open("ip-file")

	defer ip_file.Close()
	
	if isError(err) {
		return
	}

	scanner := bufio.NewScanner(ip_file)

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(),",")

		elf1 := strings.Split(assignments[0],"-")
		elf2 := strings.Split(assignments[1],"-")

		elf1_lo, _ := strconv.Atoi(elf1[0])
		elf1_hi, _ := strconv.Atoi(elf1[1])

		elf2_lo, _ := strconv.Atoi(elf2[0])
		elf2_hi, _ := strconv.Atoi(elf2[1])

		if ( ( (elf1_lo >= elf2_lo) && (elf1_hi <= elf2_hi) ) ||
			( (elf1_lo <= elf2_lo) && (elf1_hi >= elf2_hi) ) ) {

			full_overlap_count++
		}

		if ( ((elf1_lo >= elf2_lo) && (elf1_lo <= elf2_hi)) ||
			((elf1_hi >= elf2_lo) && (elf1_hi <= elf2_hi)) ||
			((elf2_lo >= elf1_lo) && (elf2_lo <= elf1_hi)) ||
			((elf2_hi >= elf1_lo) && (elf2_hi <= elf1_hi)) ) {

			partial_overlap_count++
		}
	}

	fmt.Println("Full overlap assignments = ", full_overlap_count)

	fmt.Println("Partial overlap assignments = ", partial_overlap_count)

}
