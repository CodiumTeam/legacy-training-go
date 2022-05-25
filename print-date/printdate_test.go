package print_date

import "testing"

func Test_PrintDate(t *testing.T) {
	printDate := PrintDate{
		Calendar{},
		Printer{},
	}

	printDate.printCurrentDate()

	// How can we test this function?
}
