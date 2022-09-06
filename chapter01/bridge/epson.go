// Concrete implementation
package bridge

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println(" >>> Printing by an EPSON Printer")
}
