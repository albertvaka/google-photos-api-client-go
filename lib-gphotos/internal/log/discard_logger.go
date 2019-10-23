package log

import (
	"fmt"
	"os"
)

// DiscardLogger just discards every log statement
type DiscardLogger struct {
	PanicOnExit bool
}

// Debug implements logger interface
func (d *DiscardLogger) Debug(args ...interface{}) {}

// Debugf implements logger interface
func (d *DiscardLogger) Debugf(format string, args ...interface{}) {}

// Info implements logger interface
func (d *DiscardLogger) Info(args ...interface{}) {}

// Infof implements logger interface
func (d *DiscardLogger) Infof(format string, args ...interface{}) {}

// Warn implements logger interface
func (d *DiscardLogger) Warn(args ...interface{}) {}

// Warnf implements logger interface
func (d *DiscardLogger) Warnf(format string, args ...interface{}) {}

// Error implements logger interface
func (d *DiscardLogger) Error(args ...interface{}) {}

// Errorf implements logger interface
func (d *DiscardLogger) Errorf(format string, args ...interface{}) {}

// Fatal implements logger interface
func (d *DiscardLogger) Fatal(args ...interface{}) {
	if d.PanicOnExit {
		d.Panic(args...)
	}

	os.Exit(1)
}

// Fatalf implements logger interface
func (d *DiscardLogger) Fatalf(format string, args ...interface{}) {
	if d.PanicOnExit {
		d.Panicf(format, args...)
	}

	os.Exit(1)
}

// Panic implements logger interface
func (d *DiscardLogger) Panic(args ...interface{}) {
	panic(fmt.Sprint(args...))
}

// Panicf implements logger interface
func (d *DiscardLogger) Panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

// Done implements logger interface
func (d *DiscardLogger) Done(args ...interface{}) {}

// Donef implements logger interface
func (d *DiscardLogger) Donef(format string, args ...interface{}) {}

// Fail implements logger interface
func (d *DiscardLogger) Fail(args ...interface{}) {}

// Failf implements logger interface
func (d *DiscardLogger) Failf(format string, args ...interface{}) {}
