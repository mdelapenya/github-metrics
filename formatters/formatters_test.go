package formatters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	require.Nil(t, Get(""))
	require.Nil(t, Get("foo"))
	require.Equal(t, Get("console"), ConsoleFormatter{})
}

func Test_IsValid(t *testing.T) {
	require.False(t, IsValid(""))
	require.False(t, IsValid("foo"))
	require.True(t, IsValid("console"))
}
