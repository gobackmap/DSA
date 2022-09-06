package main

import "fmt"

// powerSeries gets the power series of integer x and returns tuple of square of x
// and cube of x
func powerSeries(x int) (square int, cube int, err error) {
	square = x * x
	cube = square * x
	err = nil
	return
}

func testTuple(x int) {
	fmt.Println("\n|-----------chapter01/tuple.go-----------|")
	square, cube, err := powerSeries(x)
	fmt.Printf("   x:%d, Square:%d, Cube:%d, err:%v \n", x, square, cube, err)
	fmt.Println("|----------------------------------------|")
	fmt.Println()
}
