package formatters

import (
	"strings"

	"github.com/mdelapenya/github-metrics/types"
)

var Formatters = []string{
	"console",
}

type Formatter interface {
	Format(*types.LabelResponse)
}

func Get(formatter string) Formatter {
	if formatter == "console" {
		return ConsoleFormatter{}
	}

	return nil
}

func IsValid(formatter string) bool {
	for _, f := range Formatters {
		if strings.EqualFold(f, formatter) {
			return true
		}
	}

	return false
}
