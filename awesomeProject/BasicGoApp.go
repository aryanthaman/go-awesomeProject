package main

import (
	"fmt"
)

func add(x int, y int, z int) (int, int){
	return x + y + z, x - y - z
}

func switchTypeFunc(i interface{}){
	switch iType := i.(type) {
	case int:
		fmt.Println(i, "is of type int.")
	case bool:
		fmt.Println(i, "is of type bool.")
	case string:
		fmt.Println(i, "is of type string.")
	default:
		fmt.Printf("Other Type: %T.\n", iType)
	}
}

func mapOutput(m map[string]int){
	fmt.Println("Map:", m, "with len:", len(m))
}

//func createSlice(t type, size int, capacity int){
//	fmt.Println("Slide created.")
//}

func main() {
	a, b := add(10,10,100)
	fmt.Println(a, b)

	//Type Switch
	switchTypeFunc(false)
	switchTypeFunc(35)
	switchTypeFunc("Hello from the other side.")
	switchTypeFunc(35.5)

	//Arrays
	newArray := [5]int{1,2,3,4,5}
	fmt.Println("Array is:", newArray, "with length:",  len(newArray))
	newArray2D := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println("2D array is:", newArray2D)

	//createSlice([]int, 5, 10)

	//Slices (similar to c++ vectors)
	s := make ([]int, 5)
	fmt.Println("Initial values:", s)

	for i := 0; i < 5; i++ {
		s[i] = i+1
	}
	fmt.Println("Post initialization values:", s)

	s = append(s,6, 7, 8, 9, 10)
	fmt. Println("Final values:", s)

	//2D Slices
	slice2D := make([][]int, 5)
	for i := 0; i < 5; i++ {
		slice2D[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			slice2D[i][j] = i*i + j*j
		}
	}
	fmt.Println("Here's the 2d slide:", slice2D )

	//Maps
	m1 := make(map[string]int)
	mapOutput(m1)

	//adding Key-Val pairs to map
	m1["Joe"] = 1
	m1["Dan"] = 2
	mapOutput(m1)

	//Searching for a Key and returning value from map
	v1 := m1["Joe"]
	v2 := m1["Dan"]
	v3 := m1["John"] //returns 0 if not found.
	_,v3found := m1["John"] //returns Boolean.
	fmt.Println("values:", v1, v2, v3, v3found)

	//v3 statement can be combined as follows
	v4,v4found := m1["John"]
	fmt.Println("v4:", v4, v4found)

	//Range
	for i,num := range newArray{ // also be used as _,num
		fmt.Println("i:",i, "num:", num)
	}
	fmt.Println()
	for key,val := range m1{ // also be used as _,num
		fmt.Println(key, val)
	}

	fmt.Println()
	fmt.Println()

	//Pointers

	pointerFunc()

}
