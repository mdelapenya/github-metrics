package formatters

import (
	"strings"

	"github.com/mdelapenya/github-metrics/types"
)

var Formatters = []string{
	"console", "json",
}

type Formatter interface {
	Format(*types.LabelResponse)
}

type FileBaseFormatter struct {
	file string
}

func Get(formatter string, outputFile string) Formatter {
	if formatter == "console" {
		return ConsoleFormatter{}
	} else if formatter == "json" {
		return JsonFormatter{
			FileBaseFormatter: FileBaseFormatter{
				file: outputFile,
			},
		}
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
