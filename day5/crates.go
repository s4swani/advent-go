package main

import (
	"fmt"
// 	"strings"
 	"bufio"
 	"os"
 	"log"
)

func isError (err error) bool {

	if (err != nil) {
		log.Fatal(err)
		fmt.Println(err)
	}

	return (err != nil)
}

type crate struct {
	label uint8
	next *crate
}


func push(new_crate *crate, stack **crate) {

	if *stack != nil {
		new_crate.next = *stack
	}

	*stack = new_crate
}

func (stack *crate) pop() *crate {
	var temp *crate

	temp = stack
	stack = stack.next
	temp.next = nil

	return temp
}

func (stack *crate) print_stack(number int) {
	var temp *crate

	fmt.Printf("Stack[%d]: ", number)
	temp = stack

	for temp != nil {
		fmt.Printf("%c -> ",temp.label)
		temp = temp.next
	}
	
	fmt.Printf("nil\n")
}


func main () {
	var stack [9]**crate
	//	var temp *crate

	init_stack := [9]string{
		"DHNQTWVB",
		"DWB",
		"TSQWJC",
		"FJRNZTP",
		"GPVJMST",
		"BWFTN",
		"BLDQFHVN",
		"HPFR",
		"ZSMBLNPH",
	} 

	for i:=0; i<9; i++ {
		stack[i] = new (*crate)
		
		for j:=0; j<len(init_stack[i]); j++ {
			push(&crate{init_stack[i][j], nil}, stack[i])
		}
	}

	for i:=0; i<9; i++ {
		(*stack[i]).print_stack(i+1)
	}

	ip_file, err := os.Open("ip-file")

	defer ip_file.Close()
	
	if isError(err) {
		return
	}

	// skip the input file until instructions
	scanner := bufio.NewScanner(ip_file)
	for scanner.Scan() {

		fmt.Println (scanner.Text())
		if scanner.Text() == "" {
			break
		}
	}

	// start reading here and performing the actions


	
	//	temp = (*stack[0]).pop()
	//	fmt.Printf("Popped %c at %p\n", temp.label, &temp)
	

	

}
