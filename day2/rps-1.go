package main

import (
	"fmt"
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

func main() {

	//	const op_hand [3]string = {"A", "B", "C"}
	//	const my_hand [3]string = {"X", "Y", "Z"}

	var ip_file, error = os.Open("input-file.txt")

	var strategy string = ""

	var total_score int = 0

	scoring := make( map[string]int)
	
	scoring["A X"] = 4
	scoring["A Y"] = 8
	scoring["A Z"] = 3
	scoring["B X"] = 1
	scoring["B Y"] = 5
	scoring["B Z"] = 9
	scoring["C X"] = 7
	scoring["C Y"] = 2
	scoring["C Z"] = 6

	defer ip_file.Close()

	if isError(error) {
		return
	}
	
	scanner := bufio.NewScanner(ip_file)

	// keep reading until EOF
	for scanner.Scan() {

		strategy = scanner.Text()
		total_score += scoring[strategy]
	}

	fmt.Println("Total score = ", total_score)
}
