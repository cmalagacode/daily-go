package main

import (
	"fmt"
)
var MAX int64 = 100

func main() {
	
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	myString := fmt.Sprintf("Hi %d", 11)
	fmt.Println(myString)
	for i := 0; i < 10; i++ {
		fmt.Printf("This is iteration number # %d\n", i);
	}
	x :=  []int16 {
		3, 4, 5,
	}
	for val := range x {
		fmt.Println(val)
	}
	fmt.Println(MAX)
}