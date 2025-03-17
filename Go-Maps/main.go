package main

import(
	"fmt"
)

func main() {
	shopping := make(map[string] int)
				//OR
	//shopping := map[string] int {}

	shopping = map[string] int {
		"Apple": 10,
		"Banana": 20,
		"Orange": 30,
		"Kiwi": 40,
	}

	fmt.Println(shopping)
	
	fmt.Println(shopping["Apple"])
	fmt.Println(shopping["Banana"])
	fmt.Println(shopping["Orange"])
	fmt.Println(shopping["Kiwi"])

	// Adding a new key-value pair
	shopping["Mango"] = 50
	fmt.Println(shopping)
	
}