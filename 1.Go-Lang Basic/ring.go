package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	dataRing := ring.New(6)

	// Add data to ring using for
	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = "Data " + strconv.FormatInt(int64(i), 10) // format i to decimal
		// and set ring to next
		dataRing = dataRing.Next()
	}

	// Get all data in the ring
	dataRing.Do(func(element any) {
		fmt.Println(element)
	})
}
