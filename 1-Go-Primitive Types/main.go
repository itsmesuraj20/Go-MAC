package main

import (
	"fmt"
)

// func PrimitiveData(){
// 	// var i float32
// 	// var j float64

//Complex Data Type (Complex64 , Complex128)
// 	// var i complex64 = 1 + 2i
// 	// var j complex128 = 1 + 2i
// 	// a := 1 + 2i
// 	// b := 1 + 2i

// 	// fmt.Printf("%v,%T\n", real(a), real(a))
// 	// fmt.Printf("%v,%T\n", imag(a), imag(a))

// 	// fmt.Printf("%v,%T\n", real(b), real(b))
// 	// fmt.Printf("%v,%T\n", imag(b), imag(b))

// 	// fmt.Println("Value of i is : ", i)
// 	// fmt.Println("Value of j is : ", j)

// 	// fmt.Println("Address of i is : ", &i)
// 	// fmt.Println("Address of j is : ", &j)

// 	// fmt.Printf("%v,%T\n", i, i)
// 	// fmt.Printf("%v , %T\n", j, j)

// 	// fmt.Println(i+j)
// 	// fmt.Println(i-j)
// 	// fmt.Println(i*j)
// 	// fmt.Println(i/j)

// }

func main(){
	fmt.Println("String")


	var s string = "Hello World"
	fmt.Println(s)

	str := "Short Hello World"
	fmt.Println(str)

	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", str, str)

	var arr []byte = []byte(s)
	fmt.Printf("%v,%T\n", arr, arr)
	
	var arr1 []byte = []byte(str)
	fmt.Printf("%v,%T\n", arr1, arr1)

	
}