package apmprocessor

import (
	"context"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type spanProcessor struct {
}

func newSpanProcessor(config Config) (*spanProcessor, error) {
	return &spanProcessor{}, nil
}

func (sp *spanProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	return ptrace.Traces{}, nil
}
