// Refined abstraction
package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (mac *Mac) Print() {
	fmt.Println(" >>> Print request for mac")
	mac.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
