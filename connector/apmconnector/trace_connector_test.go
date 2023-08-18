// Copyright New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package apmconnector

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

func TestMutateOneSpan(t *testing.T) {
	traces := ptrace.NewTraces()
	resourceSpans := traces.ResourceSpans().AppendEmpty()
	resourceSpans.Resource().Attributes().PutStr("service.name", "service")
	scopeSpans := resourceSpans.ScopeSpans().AppendEmpty().Spans()
	attrs := map[string]string{
		"attrKey":      "attrValue",
		"db.statement": "select * from users",
	}
	end := time.Now()
	start := end.Add(-time.Second)
	spanValues := []TestSpan{{Start: start, End: end, Name: "span", Kind: ptrace.SpanKindServer}}
	addSpan(scopeSpans, attrs, spanValues)
	logger, _ := zap.NewDevelopment()

	MutateSpans(logger, NewSQLParser(), traces)
	dbtable, dbtablePresent := scopeSpans.At(0).Attributes().Get(DbSQLTableAttributeName)
	assert.True(t, dbtablePresent)
	assert.Equal(t, dbtable.AsString(), "users")
}
