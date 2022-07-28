package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloLzyct() {
	fmt.Println("Hello Lzycat")
}

func TestCreateGoroutine(t *testing.T) {
	// Run HelloLzyct as asynchronous
	go HelloLzyct()

	fmt.Println("Labs")

	time.Sleep(1 * time.Second)

	// Result:
	// Labs
	// Hello Lzycat
}
