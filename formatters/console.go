package formatters

import (
	"log"

	"github.com/mdelapenya/github-metrics/types"
)

type ConsoleFormatter struct{}

func (cf ConsoleFormatter) Format(lr *types.LabelResponse) {
	log.Printf("Number of Issues for %s: %d\n", lr.Label, lr.Count)
}
