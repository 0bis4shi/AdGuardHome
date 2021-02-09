package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitNext(t *testing.T) {
	const (
		a  = "a"
		b  = "b"
		c  = "c"
		sp = " "
		cl = ","
	)
	s := sp + a + cl + b + sp + cl + sp + c + sp

	for _, want := range []string{"a", "b", "c"} {
		assert.Equal(t, want, SplitNext(&s, ','))
	}
	require.Empty(t, s)
}
