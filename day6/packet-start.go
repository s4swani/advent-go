package main

import (
	"fmt"
	//	"bufio"
	"strings"
	"io/ioutil"
	"log"
)

func isError(err error) bool {

	if (err != nil) {
		log.Fatal(err)
		fmt.Println(err)
	}

	return (err != nil)
}

func not_same1(char_1 uint8, char_2 uint8) bool {
	return (char_1 != char_2)
}

// function to recursively check for dupes
func packet_start (packet string, length int) bool {
	
	if length == 0 {
		return true
	} 

	if (strings.Contains(string(packet[1:length]), string(packet[0]))) {
		return false
	} else {
		return packet_start(string(packet[1:length]), length-1)
	}
}

func get_start(ip_file []uint8, packet_length int) int {
	var i int = 0
	
	for i = 0; i < len(ip_file); i++ {
		
		if packet_start(string(ip_file[i:i + packet_length]), packet_length) {
			break
		}
	}
	
	return i + packet_length
}


func main() {
	ip_file, err := ioutil.ReadFile("ip-file")
	var packet_start, message_start int = 0, 0
	//	var charA, charB, charC, charD uint8
	
	if isError(err) {
		return
	}

	// charA = ip_file[0]
	// charB = ip_file[1]
	// charC = ip_file[2]

	// for i := 3; i < len(ip_file); i++ {
		
	// 	charD = ip_file[i]

	// 	if (not_same(charA, charB) &&
	// 		not_same(charA, charC) &&
	// 		not_same(charA, charD) &&
	// 		not_same(charB, charC) &&
	// 		not_same(charB, charD) &&
	// 		not_same(charC, charD)) {
	// 		break;
	// 	} else {
	// 		packet_start++
	// 		charA = charB
	// 		charB = charC
	// 		charC = charD
	// 	}
	// }

	packet_start = get_start(ip_file, 4)

	message_start = get_start(ip_file, 14)

	fmt.Printf("Start of packet = %d\n", packet_start)

	fmt.Printf("Start of message = %d\n", message_start)
	
}
