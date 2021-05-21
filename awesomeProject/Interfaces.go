package main

import (
	"fmt"
	"math"
)

type park interface{
	perimeter() float64
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	breadth float64
}

func (c circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64{
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64{
	return 2 * (r.length + r.breadth)
}

func printStatsForPark(p park){
	fmt.Println("area:", p.area())
	fmt.Println("perimeter:", p.perimeter())
}

func interfaceFunc(){

	fmt.Println("INTERFACES:")

	c := circle{3.0}
	r := rectangle{4.0, 5.0}
	printStatsForPark(c)
	printStatsForPark(r)

}