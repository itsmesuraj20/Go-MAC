package main

import "fmt"

func main() {

	var i int = 10
	var j float64 = 20.5

	fmt.Println("Value of i is : ", i)
	fmt.Println("Value of j is : ", j)

	fmt.Println("Address of i is : ", &i)
	fmt.Println("Address of j is : ", &j)

	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)

}