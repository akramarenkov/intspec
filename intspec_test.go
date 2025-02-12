package intspec

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeSig(t *testing.T) {
	minimum64, maximum64 := Range[int64]()
	require.Equal(t, int64(math.MinInt64), minimum64)
	require.Equal(t, int64(math.MaxInt64), maximum64)

	minimum32, maximum32 := Range[int32]()
	require.Equal(t, int32(math.MinInt32), minimum32)
	require.Equal(t, int32(math.MaxInt32), maximum32)

	minimum16, maximum16 := Range[int16]()
	require.Equal(t, int16(math.MinInt16), minimum16)
	require.Equal(t, int16(math.MaxInt16), maximum16)

	minimum8, maximum8 := Range[int8]()
	require.Equal(t, int8(math.MinInt8), minimum8)
	require.Equal(t, int8(math.MaxInt8), maximum8)

	minimum, maximum := Range[int]()
	require.Equal(t, math.MinInt, minimum)
	require.Equal(t, math.MaxInt, maximum)
}

func TestRangeUns(t *testing.T) {
	minimum64, maximum64 := Range[uint64]()
	require.Equal(t, uint64(0), minimum64)
	require.Equal(t, uint64(math.MaxUint64), maximum64)

	minimum32, maximum32 := Range[uint32]()
	require.Equal(t, uint32(0), minimum32)
	require.Equal(t, uint32(math.MaxUint32), maximum32)

	minimum16, maximum16 := Range[uint16]()
	require.Equal(t, uint16(0), minimum16)
	require.Equal(t, uint16(math.MaxUint16), maximum16)

	minimum8, maximum8 := Range[uint8]()
	require.Equal(t, uint8(0), minimum8)
	require.Equal(t, uint8(math.MaxUint8), maximum8)

	minimum, maximum := Range[uint]()
	require.Equal(t, uint(0), minimum)
	require.Equal(t, uint(math.MaxUint), maximum)
}

func TestBitSize(t *testing.T) {
	require.Equal(t, BitSize64, BitSize[int64]())
	require.Equal(t, BitSize32, BitSize[int32]())
	require.Equal(t, BitSize16, BitSize[int16]())
	require.Equal(t, BitSize8, BitSize[int8]())
	require.Equal(t, BitSizeInt, BitSize[int]())

	require.Equal(t, BitSize64, BitSize[uint64]())
	require.Equal(t, BitSize32, BitSize[uint32]())
	require.Equal(t, BitSize16, BitSize[uint16]())
	require.Equal(t, BitSize8, BitSize[uint8]())
	require.Equal(t, BitSizeInt, BitSize[uint]())
}

func BenchmarkReference(b *testing.B) {
	for b.Loop() {
		_ = b.N
	}
}

func BenchmarkRangeSig(b *testing.B) {
	for b.Loop() {
		_, _ = Range[int8]()
	}
}

func BenchmarkRangeUns(b *testing.B) {
	for b.Loop() {
		_, _ = Range[uint8]()
	}
}

func BenchmarkBitSizeSig(b *testing.B) {
	for b.Loop() {
		_ = BitSize[int8]()
	}
}

func BenchmarkBitSizeUns(b *testing.B) {
	for b.Loop() {
		_ = BitSize[uint8]()
	}
}
