package traceutil

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

const (
	// TraceKey is context trace record key
	TraceKey = "trace"
	// StartTimeKey is context start time key
	StartTimeKey = "startTime"
)

// Field is a kv pair to record additional details of the trace.
type Field struct {
	Key   string
	Value interface{}
}

func (f *Field) format() string {
	return fmt.Sprintf("%s:%v", f.Key, f.Value)
}

func writeFields(fields []Field) string {
	if len(fields) == 0 {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString("{")
	for _, f := range fields {
		buf.WriteString(f.format())
	}
	buf.WriteString("}")
	return buf.String()
}

// Trace tracing execution path
type Trace struct {
	operation    string
	lg           *zap.Logger
	fields       []Field
	startTime    time.Time
	steps        []step
	stepDisabled bool
	isEmpty      bool
}

type step struct {
	time            time.Time
	msg             string
	fields          []Field
	isSubTraceStart bool
	isSubTraceEnd   bool
}

// New return a new trace record
func New(op string, lg *zap.Logger, fields ...Field) *Trace {
	return &Trace{
		operation: op,
		lg:        lg,
		startTime: time.Now(),
		fields:    fields,
	}
}

// TODO returns a non-nil, empty Trace
func TODO() *Trace {
	return &Trace{isEmpty: true}
}

// Get return context trace record or create a new one if not exist
func Get(ctx context.Context) *Trace {
	if trace, ok := ctx.Value(TraceKey).(*Trace); ok && trace != nil {
		return trace
	}
	return TODO()
}

// GetStartTime return starttime in trace record
func (t *Trace) GetStartTime() time.Time {
	return t.startTime
}

// SetStartTime set starttime in trace record
func (t *Trace) SetStartTime(time time.Time) {
	t.startTime = time
}

// InsertStep insert one step to trace record
func (t *Trace) InsertStep(at int, time time.Time, msg string, fields ...Field) {
	newStep := step{time: time, msg: msg, fields: fields}
	if at < len(t.steps) {
		t.steps = append(t.steps[:at+1], t.steps[at:]...)
		t.steps[at] = newStep
	} else {
		t.steps = append(t.steps, newStep)
	}
}

// StartSubTrace adds step to trace as a start sign of sublevel trace
// All steps in the subtrace will logout the input fields of this function
func (t *Trace) StartSubTrace(fields ...Field) {
	t.steps = append(t.steps, step{fields: fields, isSubTraceStart: true})
}

// StopSubTrace adds step to trace as a end sign of sublevel trace
// All steps in the subtrace will logout the input fields of this function
func (t *Trace) StopSubTrace(fields ...Field) {
	t.steps = append(t.steps, step{
		fields:        fields,
		isSubTraceEnd: true,
	})
}

// Step adds step to trace
func (t *Trace) Step(msg string, fields ...Field) {
	if !t.stepDisabled {
		t.steps = append(t.steps, step{time: time.Now(), msg: msg, fields: fields})
	}
}

// StepWithFunction will measure the input function as a single step
func (t *Trace) StepWithFunction(f func(), msg string, fields ...Field) {
	t.disableStep()
	f()
	t.enableStep()
	t.Step(msg, fields...)
}

// AddFiled add fields to current trace
func (t *Trace) AddFiled(fields ...Field) {
	for _, f := range fields {
		if !t.updateFieldIfExist(f) {
			t.fields = append(t.fields, f)
		}
	}
}

// IsEmpty check whether trace is empty
func (t *Trace) IsEmpty() bool {
	return t.isEmpty
}

// Log dumps all steps in the trace
func (t *Trace) Log() {
	t.LogWithStepThreshold(0)
}

// LogIfLong dumps logs if the duration is longer than threshold
func (t *Trace) LogIfLong(threadhold time.Duration) {
	if time.Since(t.startTime) > threadhold {
		stepThreshold := threadhold / time.Duration(len(t.steps)+1)
		t.LogWithStepThreshold(stepThreshold)
	}
}

// LogWithStepThreshold only dumps step whose duration is longer than step threshold
func (t *Trace) LogWithStepThreshold(threshold time.Duration) {
	msg, fs := t.logInfo(threshold)
	if t.lg != nil {
		t.lg.Info(msg, fs...)
	}
}

func (t Trace) logInfo(threshold time.Duration) (string, []zap.Field) {
	endTime := time.Now()
	totalDuration := endTime.Sub(t.startTime)
	traceNum := rand.Int31()
	msg := fmt.Sprintf("trace[%d] %s", traceNum, t.operation)

	var steps []string
	lastStepTime := t.startTime
	for i := 0; i < len(t.steps); i++ {
		step := t.steps[i]
		// add subtrace common fields which defined at the beginning to each sub-steps
		if step.isSubTraceStart {
			for j := i + 1; j < len(t.steps) && !t.steps[j].isSubTraceEnd; j++ {
				t.steps[j].fields = append(step.fields, t.steps[j].fields...)
			}
			continue
		}

		// add subtrace common fields which defined at the end to each sub-steps
		if step.isSubTraceEnd {
			for j := i - 1; j >= 0 && !t.steps[j].isSubTraceStart; j-- {
				t.steps[j].fields = append(step.fields, t.steps[j].fields...)
			}
			continue
		}
	}

	for i := 0; i < len(t.steps); i++ {
		step := t.steps[i]
		stepDuration := step.time.Sub(lastStepTime)
		if stepDuration > threshold {
			steps = append(steps, fmt.Sprintf("trace[%d] '%v' %s (duration: %v)",
				traceNum, step.msg, writeFields(step.fields), stepDuration))
		}
		lastStepTime = step.time
	}

	fs := []zap.Field{
		zap.String("detail", writeFields(t.fields)),
		zap.Duration("duration", totalDuration),
		zap.Time("start", t.startTime),
		zap.Time("end", endTime),
		zap.Strings("steps", steps),
		zap.Int("step_count", len(steps)),
	}
	return msg, fs
}

func (t *Trace) updateFieldIfExist(f Field) bool {
	for i, v := range t.fields {
		if v.Key == f.Key {
			t.fields[i].Value = f.Value
			return true
		}
	}
	return false
}

// disableStep sets the flag to prevent the trance from adding steps
func (t *Trace) disableStep() {
	t.stepDisabled = true
}

// enableStep re-enable the trace to add steps
func (t *Trace) enableStep() {
	t.stepDisabled = false
}
