package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// Len gets the length of an IntHeap
func (ih IntHeap) Len() int {
	return len(ih)
}

// Less checks if element of i index is less than j index
func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

// Swap swaps the element of i to j index
func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

// Push pushes the item
func (ih *IntHeap) Push(item interface{}) {
	*ih = append(*ih, item.(int))
}

// Pop pops the item from the heap
func (ih *IntHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntHeap = *ih
	n = len(previous)
	x1 = previous[n-1]
	*ih = previous[0 : n-1]
	return x1
}

func testHeap() {

	fmt.Println("\n|-------------------------chapter01/heap.go-------------------------|")
	fmt.Println(" Define a IntHeap type which implements the Interface of heap package")
	fmt.Println(" Use \"container/heap\" package to init your heap.")
	fmt.Println(" Feel free to check the methods of this package.")
	fmt.Println("|-------------------------------------------------------------------|")
	fmt.Println()
	/* intHeap of type IntHeap implements the heap Intrface:
	package heap
	...
	type Interface interface {
		sort.Interface
		Push(x any) // add x as element Len()
		Pop() any   // remove and return element Len() - 1.
	}
	*/
	/* You can also see the sort package Interface:
		package sort
		...
		type Interface interface { // See the details in sort package...
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}
	*/

	intHeap := IntHeap{1, 4, 7, 8}
	heap.Init(&intHeap)
	fmt.Println("---> initial int heap:", intHeap)

	heap.Push(&intHeap, 5)
	fmt.Println("---> push 5 to int heap...", intHeap)
	heap.Push(&intHeap, 2)
	fmt.Println("---> push 2 to int heap...", intHeap)
	heap.Push(&intHeap, 9)
	fmt.Println("---> push 9 to int heap...", intHeap)

	fmt.Printf("minimum: %d\n", (intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("==> Pop %d from int heap\n", heap.Pop(&intHeap))
	}
}
