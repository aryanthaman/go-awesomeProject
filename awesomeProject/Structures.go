package main

import "fmt"


type student struct {
	name string
	age int
	quote string
}

func createStudent (name string, age int, quote string) *student{
	s3 := student{name: name}
	s3.age = age
	s3.quote = quote
	return &s3
}

func structureFunc()  {
	fmt.Println("STRUCTURES:")

	fmt.Println(student{"Alice", 15, "\"Hello!\""})

	s1 := student{"Bob", 10, "\"That's what she said.\""}
	//println(s1) //doesn't work!
	println(s1.name, s1.age, s1.quote)

	//pointer to struct
	s2 := &s1
	println(s2.name, s2.age, s2.quote) //Automatically dereferenced
	//print(*s2.name)

	//change values of struct in pointer
	s2.name = "Bobby"
	println(s2.name, s2.age, s2.quote)

	//return pointer to struct from function
	s4 := createStudent("Jason", 18, "\"Ok then.\"")
	fmt.Println(s4)
	fmt.Println(*s4)

	fmt.Println()
	fmt.Println()

	interfaceFunc()
}
