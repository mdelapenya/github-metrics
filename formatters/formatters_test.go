package formatters

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	tmp := t.TempDir()

	require.Nil(t, Get("", filepath.Join(tmp, "empty.log")))
	require.Nil(t, Get("foo", filepath.Join(tmp, "foo.log")))
	require.Equal(t, Get("console", filepath.Join(tmp, "console.log")), ConsoleFormatter{})

	jsonFile := filepath.Join(tmp, "json.log")
	require.Equal(t, Get("json", jsonFile), JsonFormatter{FileBaseFormatter{file: jsonFile}})
}

func Test_IsValid(t *testing.T) {
	require.False(t, IsValid(""))
	require.False(t, IsValid("foo"))
	require.True(t, IsValid("console"))
	require.True(t, IsValid("json"))
}
