package main

import "fmt"

// const age = 20

func main() {

	// :=
	// const name = "golang"
	// age = 20

	// fmt.Println(age)

	// grouping
	const (
		portal = 5000
		host   = "loaclhost"
	)

	//not allowed
	// portal = 3000

	fmt.Println(portal, host)

}
