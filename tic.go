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
	if len(startMsg) > 0 {
		fmt.Println(startMsg...)
	}
	start := time.Now()
	return func(endMsg ...interface{}) {
		dt := time.Now().Sub(start)
		endMsg = append(endMsg, fmt.Sprint(dt))
		fmt.Println(endMsg...)
	}
}
