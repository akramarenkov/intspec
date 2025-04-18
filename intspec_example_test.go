package intspec_test

import (
	"fmt"

	"github.com/akramarenkov/intspec"
)

func ExampleRange() {
	minimum, maximum := intspec.Range[int8]()
	fmt.Println(minimum)
	fmt.Println(maximum)
	// Output:
	// -128
	// 127
}

func ExampleBitSize() {
	size := intspec.BitSize[int8]()
	fmt.Println(size)
	// Output:
	// 8
}
