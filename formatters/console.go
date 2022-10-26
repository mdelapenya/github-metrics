package formatters

import (
	"github.com/mdelapenya/github-metrics/log"
	"github.com/mdelapenya/github-metrics/types"
	"go.uber.org/zap"
)

type ConsoleFormatter struct{}

func (cf ConsoleFormatter) Format(lr *types.MetricResponse) {
	log.Info(lr.Message, zap.Int("count", lr.Count))
}
