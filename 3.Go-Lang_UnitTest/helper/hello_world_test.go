package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldLzyct(t *testing.T) {
	result := HelloWorld("Lzyct")

	if result != "Hello Lzyct" {
		// Print the error then call t.Fail()
		t.Error("Result not match")
	}

	// This
	fmt.Println("End of function TestHelloWorldLzyct")
}

func TestHelloWorldLabs(t *testing.T) {
	result := HelloWorld("Labs")

	if result != "Hello Labs" {
		// Print the error then call t.FailNow()
		t.Fatal("Result not match")
	}

	// This not called
	fmt.Println("End of function TestHelloWorldLabs")
}
