package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)


func isError (err error) bool {

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return (err != nil)
}

func main() {

	var ip_file, error = os.Open("input-file.txt")

	var res [] string
	var rock = map[string]int  {"X":3, "Y":1, "Z":2}
	var paper = map[string]int {"X":1, "Y":2, "Z":3}
	var scissor = map[string]int {"X":2, "Y":3, "Z":1}

	var total_score int = 0

	defer ip_file.Close()

	if isError(error) {
		return
	}
	
	scanner := bufio.NewScanner(ip_file)

	// keep reading until EOF
	for scanner.Scan() {

		res = strings.SplitAfter(scanner.Text(), " ")
		opponent_hand := strings.TrimSpace(res[0])

		switch opponent_hand {			
		case "A": 
			total_score += rock[res[1]]			
		case "B":
			total_score += paper[res[1]]			
		case "C":
			total_score += scissor[res[1]]
		}

		if res[1] == "Y" {
			total_score += 3
		}

		if res[1] == "Z" {
			total_score += 6
		}

	}

	fmt.Println("Total score = ", total_score)
}
