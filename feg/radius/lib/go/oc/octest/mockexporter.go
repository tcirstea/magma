/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package octest

import (
	"sync"

	"go.opencensus.io/trace"
)

// MockExporter is only intended to test opencensus instrumentation.
type MockExporter struct {
	sync.RWMutex
	spans []*trace.SpanData
}

// ExportSpan belongs to the trace.Exporter interface.
func (e *MockExporter) ExportSpan(s *trace.SpanData) {
	e.Lock()
	e.spans = append(e.spans, s)
	e.Unlock()
}

// ExportedSpans returns all span data that have been exported since the
// MockExporter was registered or since the last call to Reset() method.
func (e *MockExporter) ExportedSpans() []*trace.SpanData {
	e.RLock()
	defer e.RUnlock()
	spans := make([]*trace.SpanData, len(e.spans))
	copy(spans, e.spans)
	return spans
}

// Reset clears the internally accumulated span data.
func (e *MockExporter) Reset() {
	e.Lock()
	e.spans = []*trace.SpanData{}
	e.Unlock()
}
