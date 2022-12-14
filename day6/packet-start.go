package main

import (
	"fmt"
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

	// can use range instead of len(ip_file)
	for i = range ip_file {
		
		if packet_start(string(ip_file[i:i + packet_length]), packet_length) {
			break
		}
	}
	
	return i + packet_length
}


func main() {
	var packet_start, message_start int = 0, 0

	const pkt_length, msg_length int = 4, 14
		
	ip_file, err := ioutil.ReadFile("ip-file")	

	if isError(err) {
		return
	}

	packet_start = get_start(ip_file, pkt_length)

	message_start = get_start(ip_file, msg_length)

	fmt.Printf("Start of packet = %d\n", packet_start)

	fmt.Printf("Start of message = %d\n", message_start)
}
