package main

import "fmt"

// const USERNAME = "root"
// const PASSWORD = "root"
const (
	i = iota+1
	j = iota
	k = iota

)

const (
	a = iota
	b
	c
)

const (
	d = iota 
	e
	_
	f
)

func main() {

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	println("-----------------")
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	println("-----------------")

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)



	// const a = 55
	// const b = 66
	// const c = a + b

	// const USERNAME = "admin"
	// const PASSWORD = "admin"


	// const i int = 12
	// const j float32 = 12.2
	// const k string = "Hello"
	// const l bool = true

	// a = 77 // Error: cannot assign to a
	// println(c)
	// fmt.Printf("%v,%T",c, c)
	// fmt.Println()
	

	// fmt.Println(USERNAME)
	// fmt.Println(PASSWORD)

	// fmt.Println(i)
	// fmt.Println(j)
	// fmt.Println(k)
	// fmt.Println(l)
	



}
