package main

import "fmt"

func endApp() {
	// catch error must in defer function
	// with recovery, to make the app run
	// even if the app crash
	message := recover()
	if message != nil {
		fmt.Println("Error :", message)
	}
	fmt.Println("Done")
}

func runApp(error bool) {
	defer endApp()
	if error {
		// throw error
		panic("Runtime exception")
	}
	fmt.Println("Application start")

}
func main() {
	runApp(true)

	// It's can't execute if you not using recover function
	fmt.Println("Hi, the app still running")
}
