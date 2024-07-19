package main

import "fmt"

func main() {

	// age := 16

	// if age == 18 {
	// 	fmt.Println("Person is adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }

	// age := 12

	// if age == 18 {
	// 	fmt.Println("Person is adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermissions = true

	// here can use OR oprator ||
	// here can use AND oprator &&

	// if role == "admin" || hasPermissions == true {
	// 	fmt.Println("has been login")
	// }

	// we can declear a veriable inside if construct
	if age := 18; age >= 18 {
		fmt.Println("Person is adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}

	// go dosen't have ternary operator , we will use normal if else
}
