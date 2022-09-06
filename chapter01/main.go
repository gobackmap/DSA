package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	target := flag.Arg(0)
	switch target {
	case "list":
		testList()
	case "tuple":
		testTuple(5)
	case "heap":
		testHeap()
	case "adapter1":
		testAdapter1()
	case "adapter2":
		testAdapter2()
	case "bridge1":
		testBridge1()
	case "bridge2":
		testBridge2()
	default:
		log.Fatal("no such test case")
	}
}
