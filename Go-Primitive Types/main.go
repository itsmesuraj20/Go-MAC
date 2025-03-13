package main

import (
	"fmt"
)

func main(){
	// var i float32
	// var j float64

	a := 1 + 2i
	b := 1 + 2i

	fmt.Printf("%v,%T\n", real(a), real(a))
	fmt.Printf("%v,%T\n", imag(a), imag(a))

	fmt.Printf("%v,%T\n", real(b), real(b))
	fmt.Printf("%v,%T\n", imag(b), imag(b))

	// fmt.Println("Value of i is : ", i)
	// fmt.Println("Value of j is : ", j)

	// fmt.Println("Address of i is : ", &i)
	// fmt.Println("Address of j is : ", &j)

	// fmt.Printf("%v,%T\n", i, i)
	// fmt.Printf("%v , %T\n", j, j)

	// fmt.Println(i+j)
	// fmt.Println(i-j)
	// fmt.Println(i*j)
	// fmt.Println(i/j)


}