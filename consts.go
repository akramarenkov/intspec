package intspec

const (
	BitSize64  = 64
	BitSize32  = 32
	BitSize16  = 16
	BitSize8   = 8
	BitSizeInt = BitSize32 << (^uint(0) >> (BitSize64 - 1))

	MaxUint64 = 1<<BitSize64 - 1
	MaxUint32 = 1<<BitSize32 - 1
	MaxUint16 = 1<<BitSize16 - 1
	MaxUint8  = 1<<BitSize8 - 1
	MaxUint   = 1<<BitSizeInt - 1

	MaxInt64 = 1<<(BitSize64-1) - 1
	MinInt64 = -1 << (BitSize64 - 1)
	MaxInt32 = 1<<(BitSize32-1) - 1
	MinInt32 = -1 << (BitSize32 - 1)
	MaxInt16 = 1<<(BitSize16-1) - 1
	MinInt16 = -1 << (BitSize16 - 1)
	MaxInt8  = 1<<(BitSize8-1) - 1
	MinInt8  = -1 << (BitSize8 - 1)
	MaxInt   = 1<<(BitSizeInt-1) - 1
	MinInt   = -1 << (BitSizeInt - 1)
)
