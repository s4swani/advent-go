package main

import (
	"fmt"
	"strings"
	 "bufio"
	 "os"
	"log"
)

func isError (err error) bool {

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return (err != nil)
}

func main () {

	/* Build the priority values map*/
	var priority = make(map[string]int)
	var total_priority int = 0

	for i:=97; i<123; i++ {
		priority[string(i)] = i-96
	}

	for i:=65; i<91; i++ {
		priority[string(i)] = i-38
	}

	var ip_file, err = os.Open("ip-file")

	defer ip_file.Close()
	
	if isError(err) {
		return
	}

	scanner := bufio.NewScanner(ip_file)

	/* go through each line (rucksack) in the input file */
	for scanner.Scan() {
		rucksack := scanner.Text()
		length := len(rucksack)
		
		/* split the line in two */
		firstComp := rucksack[0:length/2]
		secondComp := rucksack[length/2:length]

		/*
		 * go through each character in the first half
		 * and check if it exists in the second half 
		 */

		for i:=0; i<len(firstComp); i++ {
			if strings.Contains(secondComp, string(firstComp[i])) {
				total_priority += priority[string(firstComp[i])]
				break //only need to count once
			}
		}
	}

	fmt.Println("Total priority = ", total_priority)


}
