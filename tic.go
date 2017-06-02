/*
Package tic provides an easy way to time how long a function runs.

Here is an example of how to use it:

	func main() {
		defer tic.Toc("start timer")("work took:")
		fmt.Println("this takes some time...")
		// ... (do some actual work here)
	}

The above program will print something like:
	start timer
	this takes some time...
	work took: 2.6973829s

Start and end messages are optional so the program:

	func main() {
		defer tic.Toc()()
	}

will only print the time and nothing else.

You can change the Println function to whatever you want the output to be, if
you set it to nil there will be no output. By default it is set to fmt.Println.
Println is captured when calling Toc, this means it will use the value that
Println had when calling Toc, even if you set Println to something else before
the deferred timer function runs.
*/
package tic

import (
	"fmt"
	"time"
)

// Toc returns a function that prints the time between the call to Toc and the
// call to the returned function. See the package documentation for further
// information.
func Toc(startMsg ...interface{}) func(endMsg ...interface{}) {
	pl := Println
	if pl == nil {
		pl = dummyPrintln
	}
	if len(startMsg) > 0 {
		pl(startMsg...)
	}
	start := time.Now()
	return func(endMsg ...interface{}) {
		dt := time.Now().Sub(start)
		endMsg = append(endMsg, fmt.Sprint(dt))
		pl(endMsg...)
	}
}

var Println = fmt.Println

func dummyPrintln(a ...interface{}) (int, error) { return 0, nil }
