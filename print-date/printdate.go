package print_date

import (
	"fmt"
	"time"
)

type PrintDate struct {
	calendar Calendar
	printer  Printer
}

func (printDate PrintDate) printCurrentDate() {
	today := printDate.calendar.today()
	printDate.printer.printLine(today.String())
}

type Calendar struct {
}

func (calendar Calendar) today() time.Time {
	return time.Now()
}

type Printer struct {
}

func (printer Printer) printLine(line string) {
	fmt.Println(line)
}
