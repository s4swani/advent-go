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


	var foundInThree bool
	
	var badgeToken string = ""

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
		
		for i:=0; i<len(rucksackA); i++ {
			
			badgeToken = ""
			
			foundInThree = false

			if strings.Contains(rucksackB, string(rucksackA[i])) {
				
				for j:=0; j<len(rucksackC); j++ {
					if strings.Contains(rucksackC, string(rucksackA[i])) {
						foundInThree = true
						badgeToken = string(rucksackA[i])
						break
					}
				}
				
				if foundInThree {
					break
				}				
			}
		}
	
		if foundInThree {
			total_priority += priority[badgeToken]
		}
	}

	fmt.Println ("Total priority = ", total_priority)
}

