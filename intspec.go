// Library that provides to detect the parameters of integer types: minimum, maximum
// values and bit size.
package intspec

import (
	"golang.org/x/exp/constraints"
)

// Returns minimum and maximum values for specified integer type.
func Range[Type constraints.Integer]() (Type, Type) {
	if !isSigned[Type]() {
		return 0, ^Type(0) // Type(0)-1
	}

	refMaxInt64 := int64(MaxInt64)
	refMinInt64 := int64(MinInt64)
	refMaxInt32 := int32(MaxInt32)
	refMinInt32 := int32(MinInt32)
	refMaxInt16 := int16(MaxInt16)
	refMinInt16 := int16(MinInt16)
	refMaxInt8 := int8(MaxInt8)
	refMinInt8 := int8(MinInt8)

	convMaxInt64 := Type(refMaxInt64)
	convMinInt64 := Type(refMinInt64)
	convMaxInt32 := Type(refMaxInt32)
	convMinInt32 := Type(refMinInt32)
	convMaxInt16 := Type(refMaxInt16)
	convMinInt16 := Type(refMinInt16)
	convMaxInt8 := Type(refMaxInt8)
	convMinInt8 := Type(refMinInt8)

	switch {
	case int64(convMaxInt64) == refMaxInt64:
		return convMinInt64, convMaxInt64
	case int32(convMaxInt32) == refMaxInt32:
		return convMinInt32, convMaxInt32
	case int16(convMaxInt16) == refMaxInt16:
		return convMinInt16, convMaxInt16
	}

	return convMinInt8, convMaxInt8
}

// Returns bit size for specified integer type.
func BitSize[Type constraints.Integer]() int {
	refMaxInt64 := int64(MaxInt64)
	refMaxInt32 := int32(MaxInt32)
	refMaxInt16 := int16(MaxInt16)

	convMaxInt64 := Type(refMaxInt64)
	convMaxInt32 := Type(refMaxInt32)
	convMaxInt16 := Type(refMaxInt16)

	refMaxUint64 := uint64(MaxUint64)
	refMaxUint32 := uint32(MaxUint32)
	refMaxUint16 := uint16(MaxUint16)

	convMaxUint64 := Type(refMaxUint64)
	convMaxUint32 := Type(refMaxUint32)
	convMaxUint16 := Type(refMaxUint16)

	if isSigned[Type]() {
		switch {
		case int64(convMaxInt64) == refMaxInt64:
			return BitSize64
		case int32(convMaxInt32) == refMaxInt32:
			return BitSize32
		case int16(convMaxInt16) == refMaxInt16:
			return BitSize16
		}

		return BitSize8
	}

	switch {
	case uint64(convMaxUint64) == refMaxUint64:
		return BitSize64
	case uint32(convMaxUint32) == refMaxUint32:
		return BitSize32
	case uint16(convMaxUint16) == refMaxUint16:
		return BitSize16
	}

	return BitSize8
}
