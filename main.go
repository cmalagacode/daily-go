package main

import (
	"fmt"
	"math"
	"errors"
	"strconv"
	"os"
)
const MAX int64 = 100

func Add(x int, y int) (int, error) {
	result := x + y
	if result < 10 {
		return result, errors.New("less than threshold")
	}
	return result, nil
}

func VariablePract() {
	i := 1
	for i < 201 {
		fmt.Printf("This is iteration number # %v\n", i)
		i++
	}
	fmt.Println("End of the road")
}

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
	fmt.Println("============")
	const PI float64 = 3.14
	fmt.Println(PI * 1000)
	fmt.Println("============")
	fmt.Println("============")
	y := 12.32
	fmt.Printf("Full control of text %v\n", y)
	fmt.Println("============")
	fmt.Println("============")
	a := 5.048
	b := 10.0230
	k := math.Abs(a - b)
	fmt.Printf("%.2f\n", k)
	fmt.Println("============")

	result, err := Add(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	ch := 'E'
	switch (ch) {
	case 'A':
		fmt.Println("foo")
	case 'K':
		fmt.Println("bar")
	default:
		fmt.Println("All else is at work")
	}
	lenArgs := len(os.Args)
	fmt.Println(lenArgs)
	fmt.Println("-======")
	v := "22"
	newV, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println("Stop here")
	}
	fmt.Println(newV)
	fmt.Println("-======")
	fmt.Println("-======")
	VariablePract()
	fmt.Println("-======")
}