package formatters

import (
	"github.com/mdelapenya/github-metrics/log"
	"github.com/mdelapenya/github-metrics/types"
	"go.uber.org/zap"
)

type ConsoleFormatter struct{}

func (cf ConsoleFormatter) Format(lr *types.MetricResponse) {
	fields := []log.Field{
		zap.Int("count", lr.Count),
		zap.String("query", lr.Query),
	}

	for k, v := range lr.Metadata {
		fields = append(fields, zap.String(k, v))
	}

	log.Info(lr.Title, fields...)
}
