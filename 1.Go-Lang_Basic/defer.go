package main

import "fmt"

func logging() {
	fmt.Println("Done")
}

func startApplication(value int) {
	// execute function after startApplication
	// even the startApplication have errors
	defer logging()

	fmt.Println("Application Start ...")

	result := 10 / value
	fmt.Println("Result", result)

}

func main() {

	startApplication(10)
	// Application Start ...
	// Result 1
	// Done

	startApplication(0)
	// Application Start ...
	// Done
	// panic: runtime error: integer divide by zero
}
