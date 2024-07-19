package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other values")
	// }

	//mulitple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's work day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float364")
		default:
			fmt.Println("unknow", t)
		}
	}
	whoAmI(false)
}
