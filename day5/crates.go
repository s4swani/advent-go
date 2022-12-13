package main

import (
	"fmt"
// 	"strings"
// 	"bufio"
// 	"os"
// 	"log"
)

// func isError (err error) bool {

// 	if (err != nil) {
// 		log.Fatal(err)
// 		fmt.Println(err)
// 	}

// 	return (err != nil)
// }

type crate struct {
	label uint8
	next *crate
}


func push(new_crate *crate, stack **crate) {

	fmt.Printf ("\n%p %p\n", new_crate, *stack)
	
	if (*stack) != nil {
		//		fmt.Printf("adding element: %c\n",new_crate.label)
		new_crate.next = *stack
	}
	
	*stack = new_crate
	fmt.Printf ("%p%v %p%v\n\n", new_crate, new_crate, *stack, **stack)
}

// func (stack *crate) pop() *crate {

// 	temp := stack
// 	stack = stack.next
// 	temp.next = nil

// 	return &temp
// }

func (stack *crate) print_stack() {
	var temp *crate
	temp = stack
	
	for temp != nil {
		fmt.Printf("%c -> ",temp.label)
				
		temp = temp.next
	}
	fmt.Printf("nil\n")
}


func main () {
	var stacks = new(*crate)
	var temp crate = crate{'A',nil}
		

	//	fmt.Printf("%p%v %p%v\n", &temp,temp, stacks,*stacks)

	push(&temp, stacks)

	//	fmt.Printf("%p%v %p%v\n", &temp,temp, stacks,*stacks)
	
	push(&crate{'B', nil}, stacks)

	//	fmt.Printf("%p%v %p%v\n", &temp,temp, stacks,*stacks)

	push(&crate{'C', nil}, stacks)

	push(&crate{'D', nil}, stacks)

	(*stacks).print_stack()


}
