package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var f0 int = 0
	var f1 int = 1
	var result int = 0
	return func() int {

		if result == 0 {
			fmt.Print(result, ", ", f1, ", ")
		}

		result = f0 + f1
		f0 = f1
		f1 = result

		return result
	}
}
func main() {
	f := fibonacci()
	defer fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Print(f())
		} else {
			fmt.Print(f(), ", ")
		}
	}
}
