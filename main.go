package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Provided to many args")
		return
	}
	for _, arg := range os.Args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Not a number")
			continue
		}
		if num == 10 {
			fmt.Println("Foo")
		} else if num > 10 {
			fmt.Println("foooo")
		} else {
			fmt.Println("oooo")
		}
	}
}