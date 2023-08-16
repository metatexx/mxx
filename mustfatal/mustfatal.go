package mustfatal

import (
	"log"
	"log/slog"
	"os"
	"runtime"
)

// SLog is set to use the slog package (skips any log logger)
var SLog *slog.Logger

var Log *log.Logger

//var LogSLog slog.Logger

// Ok checks for error and exits the program if not nil
func Ok(err error) {
	FatalX(err, 1, "")
}

// Ignore ignores the error code
func Ignore(_ error) {
}

// OkSkipOne error must be not nil and skips one additional return value
func OkSkipOne[T any](_ T, err error) {
	FatalX(err, 1, "")
}

// OkSkipTwo error must be not nil and skips two additional return values
func OkSkipTwo[T any, S any](_ T, _ S, err error) {
	FatalX(err, 1, "")
}

// OkOne error must be not nil and returns one value
func OkOne[T any](arg T, err error) T {
	FatalX(err, 1, "")
	return arg
}

// IgnoreOne error must be not nil and returns one value
func IgnoreOne[T any](arg T, _ error) T {
	return arg
}

// OkTwo error must be not nil and returns two values
func OkTwo[T any, S any](arg1 T, arg2 S, err error) (T, S) {
	FatalX(err, 1, "")
	return arg1, arg2
}

// Fatal is like "Ok" but can print a msg.
func Fatal(err error, msg string) {
	FatalX(err, 1, msg)
}

// FatalX when err != nil it prints msg, then file and line of the "skip" (usually 0) caller and the panics
func FatalX(err error, skip int, msg string) {
	if err == nil {
		return
	}
	fpcs := make([]uintptr, 1)

	// Skip 2 levels to get the caller
	n := runtime.Callers(skip+2, fpcs)
	if n == 0 {
		panic("no caller")
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		panic("caller is nil")
	}

	// Print the name of the function (but trim to our package names)
	file, line := caller.FileLine(fpcs[0] - 1)
	if SLog != nil {
		if msg == "" {
			msg = "fatal error"
		}
		slog.Error(msg,
			slog.String("err", err.Error()),
			slog.Group("caller", slog.String("file", file), slog.Int("line", line)))
	}
	if Log == nil {
		if SLog != nil {
			// only use SLog logger
			os.Exit(5)
		}
		Log = log.Default()
	}
	if msg != "" {
		Log.Println("CAUSE: " + msg)
	}
	Log.Printf("CALLER: %s:%d", file, line)
	Log.Printf("FATAL: %s", err)
	os.Exit(5)
}
