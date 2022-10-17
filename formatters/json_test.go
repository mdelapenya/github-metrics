package formatters

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mdelapenya/github-metrics/types"
	"github.com/stretchr/testify/require"
)

func TestFormatJson_One(t *testing.T) {
	tmp := t.TempDir()
	defer func() {
		data = make([]*types.LabelResponse, 0)
	}()

	formatter := Get("json", filepath.Join(tmp, "output_one.json"))
	formatter.Format(&types.LabelResponse{Label: "label", Count: 1})

	labels, err := formatter.(JsonFormatter).read()
	require.Nil(t, err)
	require.Equal(t, 1, len(labels))
}

func TestFormatJson_Multiple(t *testing.T) {
	tmp := t.TempDir()
	defer func() {
		data = make([]*types.LabelResponse, 0)
	}()

	length := 125

	formatter := Get("json", filepath.Join(tmp, "output_multiple.json"))

	for i := 0; i < length; i++ {
		lr := &types.LabelResponse{Label: fmt.Sprintf("label_%d", i), Count: i}
		formatter.Format(lr)
	}

	labels, err := formatter.(JsonFormatter).read()
	require.Nil(t, err)
	require.Equal(t, length, len(labels))
}
