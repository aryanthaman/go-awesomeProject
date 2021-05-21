package main

import "fmt"

func passByValue(val int){
	val = 50
}

func passByPtr(ptr *int){
	*ptr = 50
}

func pointerFunc(){

	fmt.Println("POINTERS:")

	a := 1
	fmt.Println("a:", a)

	b := &a //b is pointer to a
	fmt.Println("b:", b)

	//dereferencing b
	fmt.Println("&b:", &b)

	//dereferencing b
	fmt.Println("*b:", *b)

	x := &a
	fmt.Println("x:", x)

	c := 5
	b = &c
	fmt.Println("new *b:", *b)



	passByValue(a)
	fmt.Println("a (after passByValue):", a)

	passByPtr(&a) //by passing &a, ptr becomes a pointer to a.
	fmt.Println("a (after passByPtr):", a)

	fmt.Println()
	fmt.Println()

	structureFunc()
}
