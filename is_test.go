package intspec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSigned(t *testing.T) {
	require.True(t, isSigned[int]())
	require.False(t, isSigned[uint]())
}

func BenchmarkIsSigned(b *testing.B) {
	for b.Loop() {
		_ = isSigned[int]()
	}
}
