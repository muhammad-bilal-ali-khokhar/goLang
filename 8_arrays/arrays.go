package main

import "fmt"

//number sequance of specific lenght
func main() {

	// here are zeroed values
	//int -> 0, string -> "", bool -> false

	// define array len with type
	var nums [4]int
	// how to check lenght of array
	checkLenght := len(nums)
	// how to push the item in array
	nums[0] = 1
	// how to get the item in array
	getValue := nums[0]
	fmt.Println(nums, getValue, checkLenght)
	//result [1 0 0 0] 1 4

	// fmt.Println(nums, checkLenght)
	// result is [1 0 0 0] 4 because of type other items 0

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)
	//resut [false false true false]

	var vals [4]string
	vals[2] = "muhammad bilal ali khokhar"
	fmt.Println(len(vals), vals)
	//resut 4 [  muhammad bilal ali khokhar ]

	//to declear it in single line
	num := [3]int{1, 2, 3}
	fmt.Println(num)

	// 2dimensional arrays
	twoDimensional := [2][2]int{{3, 4}, {6, 7}}

	fmt.Println(twoDimensional)

	//array use case
	//- fixed size , that is pedictable
	//- memory optimazation
	//- constant time access , fast access

}
