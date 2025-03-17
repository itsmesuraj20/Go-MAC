package main

import(
	"fmt"
)

func main() {
	// var arr[5]int = [5]int {1,2,3,4,5}
	// amt:= [5]float64{1.0,2.0,3.0,4.0,5.0}

	// l1 := len(arr)
	// l2 := len(amt)

	// fmt.Printf("%v,%T\n", amt, amt)
	// fmt.Print(amt)
	// fmt.Println()

	// fmt.Printf("%v,%T\n", arr, arr)
	// fmt.Print(arr)
	// fmt.Println()

	// fmt.Println("Len 1 :", l1)
	// fmt.Println("Len 2 :", l2)

	// // If the size of the array is not mentioned explicitly, it will be inferred from the number of elements
	// arr1 := [...]int{100,211,311,422,335}

	// fmt.Printf("%v,%T\n", arr1, arr1)
	// fmt.Print(arr1)
	// fmt.Println()



	// // Slices
	// newArr := []int {15,30,45,60,75}
	// fmt.Println("newARR :",newArr)

	// newArr2 := &newArr
	// fmt.Println("newARR2:", *newArr2)

	// // Slice of a slice
	// newSlice1 := newArr[:] // Slice of the entire array
	// newSlice2 := newArr[1:3] // 1 is inclusive and 3 is exclusive
	// newSlice3 := newArr[:3] //	0 is inclusive and 3 is exclusive
	// newSlice4 := newArr[2:] // 2 is inclusive and till the end
	// newSlice5 := newArr[0:4] // 1 is inclusive and 4 is exclusive
	// 		//OR
	// newSlice5v2 := newArr[:len(newArr)-1] // 1 is inclusive and 5 is exclusive		
	// newSlice6 := newArr[1:] // 1 is inclusive and till the end

	// fmt.Println()
	// fmt.Println("--------------------")
	// fmt.Println()
	// fmt.Println("newSlice1:",newSlice1)
	// fmt.Println("newSlice2:",newSlice2)
	// fmt.Println("newSlice3:",newSlice3)
	// fmt.Println("newSlice4:",newSlice4)
	// fmt.Println("newSlice5:",newSlice5)
	// fmt.Println("-----OR-----")
	// fmt.Println("newSlice5v2:",newSlice5v2)
	// fmt.Println("newSlice6:",newSlice6)

	//Slice Append

	var a []int = []int {1,2,3}
	var b []int = append(a,5)    //o/p - [1 2 3 5]
	var c []int = append(b,a...) //o/p- [1 2 3 5 1 2 3]  
	var d []int = append(a,b...) //o/p- [1 2 3 1 2 3 5]
	var e []int = append(a,a...) //o/p- [1 2 3 1 2 3]
	var f []int = append(a,c...) //o/p- [1 2 3 1 2 3 5 1 2 3]
	var g []int = append(a,d...) //o/p- [1 2 3 1 2 3 1 2 3 5]


	fmt.Println("a:",a)
	fmt.Println("b:",b)
	fmt.Println("c:",c)
	fmt.Println("d:",d)
	fmt.Println("e:",e)
	fmt.Println("f:",f)
	fmt.Println("g:",g)


}
