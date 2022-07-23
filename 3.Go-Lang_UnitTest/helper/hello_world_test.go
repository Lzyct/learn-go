package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Lzyct")

	// Using Fail
	assert.Equal(t, "Hello Lzyct", result, " The result must be Hello Lzyct")

	// This called
	fmt.Println("End of function TestHelloWorldLzyct")
}

func TestHelloWorldLzyctAssert(t *testing.T) {
	result := HelloWorld("Lzyct")

	// Using Fail
	assert.Equal(t, "Hello Lzyct", result, " The result must be Hello Lzyct")

	// This called
	fmt.Println("End of function TestHelloWorldLzyct")
}

func TestHelloWorldLabsRequire(t *testing.T) {
	result := HelloWorld("Labs")

	// Using FailNow
	require.Equal(t, "Hello Labs", result, "The result must be Hello Labs")

	// This not called
	fmt.Println("End of function TestHelloWorldLabs")
}

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
