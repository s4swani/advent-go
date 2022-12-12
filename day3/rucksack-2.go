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

	var priority = make(map[string]int)
	var total_priority int = 0
	var badgeToken string = ""

	/* Build the priority map */
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

	for scanner.Scan() {

		rucksackA := scanner.Text()
		
		scanner.Scan()
		rucksackB := scanner.Text()
		
		scanner.Scan()
		rucksackC := scanner.Text()
		
		for i, _ := range rucksackA {
			
			badgeToken = string(rucksackA[i])
	
			if strings.Contains(rucksackB, badgeToken) &&
				strings.Contains(rucksackC, badgeToken) {
				
				total_priority += priority[badgeToken]	
				break
			}
		}
	}
	fmt.Println ("Total priority = ", total_priority)
}

