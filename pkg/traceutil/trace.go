package traceutil

import "go.uber.org/zap"

type Field struct {
	Key   string
	Value interface{}
}

type Trace struct {
	operation string
	lg        *zap.Logger
	fields    []Field
}
