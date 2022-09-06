package main

import (
	"container/list"
	"fmt"
)

func testList() {
	fmt.Println("\n|-----------------chapter01/list.go-----------------|")
	fmt.Println(" Use \"container/list\" package to create a new list.")
	fmt.Println(" Feel free to check the methods of this package.")
	fmt.Println("|---------------------------------------------------|")
	fmt.Println()

	var intList list.List
	intList.PushBack(11)
	ele23 := intList.PushBack(23)
	intList.PushBack(45)
	intList.InsertAfter(34, ele23)

	fmt.Printf("Int List: ")
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d", element.Value)
		if element.Next() != nil {
			fmt.Printf("<=>")
		}
	}
	fmt.Printf("\n\n")
}
