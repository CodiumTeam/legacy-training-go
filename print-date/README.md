# Goal
Be able to test printCurrentDate function without changing the method signature.

1. Test the code with doubles using a library.
2. Test the code with doubles created by you.
# Code to test
    func (printDate PrintDate) printCurrentDate() {
        today := printDate.calendar.today()
        printDate.printer.printLine(today.String())
    }
# Learnings

How to use TBD to generate the doubles.

How to build a Mock and Stub manually.

## Tools

### Example of spy

TBD
	
### Example of stub

TBD


## Authors
Luis Rovirosa [@luisrovirosa](https://www.twitter.com/luisrovirosa)

Jordi Anguela [@jordianguela](https://www.twitter.com/jordianguela)