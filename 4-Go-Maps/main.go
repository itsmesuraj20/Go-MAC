package main

import (
	"fmt"
)

func main() {
	shopping := make(map[string] int)
				//OR
	//shopping := map[string] int {}

	shopping = map[string] int {
		"Apple" : 10,
		"Banana": 20,
		"Suraj" : 30,
		"Orange": 30,
		"Kiwi"  : 40,
	}

	fmt.Println(shopping)
	fmt.Println(len(shopping))
	fmt.Println(reflect.TypeOf(shopping))
	
	// fmt.Println(shopping["Apple"])
	// fmt.Println(shopping["Banana"])
	// fmt.Println(shopping["Orange"])
	// fmt.Println(shopping["Kiwi"])

	// // Adding a new key-value pair
	// shopping["Mango"] = 50
	// fmt.Println(shopping)


	// fmt.Println(" ---------- ------------")
	// //check if a key exists
	
	// _, ok := shopping["Apple"] // _ is a blank identifier , that will ignore the value of the key 
	// fmt.Println(ok)
	// 			//OR

	// cart, ok := shopping["Loan"]
	// fmt.Println(cart, ok)
	
	sc := shopping
	fmt.Println(shopping)
	fmt.Println(sc)
	
	sc["Suraj"] = 99999999 //while copying do consider that it will change the original map as well , as it is a reference type 
	
	fmt.Println(shopping)
	fmt.Println(sc)
	
}