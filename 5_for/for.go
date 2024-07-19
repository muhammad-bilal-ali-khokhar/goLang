package main

import "fmt"

// for -> only construct in go for looping
func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinit loop
	// for {
	// 	println("infinit loop")
	// }

	// classic for loop
	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	// classic for loop break and continue
	// for i := 0; i <= 3; i++ {
	// 	// break

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//go new feat 1.22 range
	for i := range 11 {
		fmt.Println(i)
	}

}
