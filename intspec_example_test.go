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
	bisize := intspec.BitSize[int8]()
	fmt.Println(bisize)
	// Output:
	// 8
}
