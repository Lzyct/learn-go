package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Lzyct",
			request:  "Lzyct",
			expected: "Hello Lzyct",
		},
		{
			name:     "Labs",
			request:  "Labs",
			expected: "Hello Labs",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			// Using FailNow
			require.Equal(t, test.expected, result, "The result must be "+test.expected)

		})
	}
}

func TestSubTest(t *testing.T) {
	// To run specific sub unit test
	// go test -v -run=MainFunction/SubFunction
	// go test -v -run=TestSubTest/Lzyct
	t.Run("Lzyct", func(t *testing.T) {
		result := HelloWorld("Lzyct")

		// Using Fail
		assert.Equal(t, "Hello Lzyct", result, " The result must be Hello Lzyct")

	})

	t.Run("Labs", func(t *testing.T) {
		result := HelloWorld("Labs")

		// Using FailNow
		require.Equal(t, "Hello Labs", result, "The result must be Hello Labs")

	})
}

// Only work on current package
func TestMain(m *testing.M) {
	println("Initialize script before run Unit Test")
	m.Run()
	println("Clean up after unit test")
}

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
