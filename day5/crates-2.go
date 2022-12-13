package main

import (
	"fmt"
 	"strings"
 	"bufio"
 	"os"
 	"log"
	"strconv"
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

func pop(stack **crate) *crate {
	var temp *crate = nil

	if *stack != nil {

		temp = *stack
		*stack = (*stack).next
		temp.next = nil
	}

	return temp
}

func pushN(head *crate, tail *crate, stack **crate) {

	tail.next = *stack
	*stack = head
	
}

func popN(stack **crate, number int) (*crate, *crate) {
	var head, tail *crate = nil, nil

	head = *stack

	for i := 0; i < number; i++ {
		tail = *stack
		*stack = (*stack).next
	}
	
	return head, tail
}

func (stack *crate) print_stack(number int) {
	var temp *crate

	fmt.Printf("Stack[%d]: ", number)
	temp = stack

	for temp != nil {
		fmt.Printf("%c -> ", temp.label)
		temp = temp.next
	}
	
	fmt.Printf("nil\n")
}


func main () {
	var stack [9]**crate
	var head, tail *crate 

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

	// initialise the stacks
	for i:=0; i<9; i++ {
		stack[i] = new (*crate)
		
		for j:=0; j<len(init_stack[i]); j++ {
			push(&crate{init_stack[i][j], nil}, stack[i])
		}
	}

	ip_file, err := os.Open("ip-file")

	defer ip_file.Close()
	
	if isError(err) {
		return
	}

	// skip the input file until instructions
	scanner := bufio.NewScanner(ip_file)
	for scanner.Scan() {

		if scanner.Text() == "" {
			break
		}
	}

	// start reading here and performing the actions
	for scanner.Scan() {
		
		actions := strings.Split(scanner.Text(), " ")

		how_many, _ := strconv.Atoi(actions[1])
		from_index, _ := strconv.Atoi(actions[3])
		to_index, _ := strconv.Atoi(actions[5])

		from_index -= 1
		to_index -= 1

		head, tail = popN(stack[from_index], how_many)
		pushN(head, tail, stack[to_index])
	}

	for i:=0; i<9; i++ {
		(*stack[i]).print_stack(i+1)
	}
}
