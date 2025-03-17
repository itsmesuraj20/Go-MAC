package main

import "fmt"

func main(){
	var matrix [3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	matrix[0][1] = 0
	matrix[0][2] = 0
	matrix[1][0] = 0
	matrix[1][1] = 1
	matrix[1][2] = 0
	matrix[2][0] = 0
	matrix[2][1] = 0
	matrix[2][2] = 1
	fmt.Println(matrix)
}