package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Number ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

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
	// Hello Lzyc
}
