tic
===

Package tic provides an easy way to time how long a function runs.
Here is an example of how to use it:
```Go
func main() {
	defer tic.Toc("start timer")("work took:")
	fmt.Println("this takes some time...")
	// ... (do some actual work here)
}
```
The above program will print something like:
```
	start timer
	this takes some time...
	work took: 2.6973829s
```
Start and end messages are optional so the program:
```Go
func main() {
	defer tic.Toc()()
}
```
will only print the time and nothing else.
