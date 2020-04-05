package introduction

import (
	"fmt"
)

/*
DataRace ...
There are 3 likely outcome of the following function:
* Nothing is printed. In this case, go func() executed before caller function
• “the value is 0” is printed. In this case, caller function executed before go func()
• “the value is 1” is printed. In this case, "B" was executed before "A", but "A"
was executed before "C".
*/
func DataRace() {
	data := 0
	go func() {
		data++ // A is crticaal section
	}()
	if data == 0 { // B is crticaal section
		fmt.Printf("the value is %v.\n", data) // C is crticaal section
	}
}

/*
Critical Section: section of your program that needs exclusive access to a
shared resource
*/
