package main

import (
	"fmt"

	"github.com/goplateau/DSA/chapter01/bridge"
)

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct {
}

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Printf(" >>> drawing shape with x=%v, and y=%v\n", x, y)
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// DrawContour method drawContour given the coordinates
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println(" >>> drawing contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour *DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func testBridge1() {
	fmt.Println("\n|-------------------------chapter01/bridge.go-------------------------|")
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = &DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
	fmt.Println("|---------------------------------------------------------------------|")
	fmt.Println()
}

// Client code
func testBridge2() {
	fmt.Println("\n|-------------------------------chapter01/bridge/*-------------------------------|")
	fmt.Println(" - Two types of computers: Mac and Windows, and")
	fmt.Println(" - Two types of printers: Epson and HP")
	fmt.Println(" - Instead of creating 4 structs for the 2*2 combination, we create 2 hierarchies:")
	fmt.Println(" 		> Abstraction hierarchy: this will be our computers")
	fmt.Println(" 		> Implementation hierarchy: this will be our printers")
	fmt.Println(" - They communicate with each other via a Bridge, where the Abstraction (computer)")
	fmt.Println("   contains a reference to the Implementation (printer)")
	fmt.Println("|--------------------------------------------------------------------------------|")
	fmt.Println()

	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridge.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
