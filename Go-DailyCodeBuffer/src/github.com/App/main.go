package main

import "fmt"

func main() {

	// var i int = 10
	// var j float64 = 20.5

	//Type Conversion
	var a = 10
	var b float32 = float32(a)

	fmt.Println("Value of a is : ", a)
	fmt.Println("Value of b is : ", b)

	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)

	//fmt.Println("Value of i is : ", i)
	//fmt.Println("Value of j is : ", j)

	//fmt.Println("Address of i is : ", &i)
	//fmt.Println("Address of j is : ", &j)

	//fmt.Printf("%v,%T\n", i, i)
	//fmt.Printf("%v , %T\n", j, j)

}