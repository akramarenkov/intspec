package intspec

import "golang.org/x/exp/constraints"

// Detects whether the integer type is signed or unsigned.
func isSigned[Type constraints.Integer]() bool {
	number := Type(0)

	number--

	return number < 0
}
