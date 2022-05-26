package printdate

import (
	"fmt"
	"time"
)

type PrintDate interface {
	PrintCurrentDate()
}

func NewPrintDate(calendar Calendar, printer Printer) PrintDate {
	return printDate{calendar: calendar, printer: printer}
}

type printDate struct {
	calendar Calendar
	printer  Printer
}

func (printDate printDate) PrintCurrentDate() {
	today := printDate.calendar.today()
	printDate.printer.printLine(today.String())
}

type Calendar interface {
	today() time.Time
}

func NewCalendar() Calendar {
	return calendar{}
}

type calendar struct {
}

func (calendar calendar) today() time.Time {
	return time.Now()
}

type Printer interface {
	printLine(line string)
}

func NewPrinter() Printer {
	return printer{}
}

type printer struct {
}

func (printer printer) printLine(line string) {
	fmt.Println(line)
}
