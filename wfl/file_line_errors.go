package wfl

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Error wraps the error message and puts file and line in the front of it.
// Notice: If error is nil it will just return nil
func Error(err error) error {
	return ErrorWithSkip(err, 2)
}

// ErrorWithSkip wraps the error message and puts file and line in the front of it.
// The argument skip is the number of stack frames to ascend
// Notice: If error is nil it will just return nil
func ErrorWithSkip(err error, skip int) error {
	if err == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Errorf("%s:%d: %w", filepath.Base(file), line, err)
}

// Errorf wraps the Errorf output and puts file and line in the front of it.
// Notice: If there are error types in the arguments and all of them are nil,
// the function will nil
func Errorf(format string, args ...any) error {
	return ErrorfWithSkip(2, format, args)
}

// ErrorfWithSkip wraps the Errorf output and puts file and line in the front of it.
// The argument skip is the number of stack frames to ascend
// Notice: If there are error types in the arguments and all of them are nil,
// the function will nil
func ErrorfWithSkip(skip int, format string, args ...any) error {
	_, file, line, _ := runtime.Caller(skip)
	format = "%s:%d: " + format
	nargs := make([]any, 0, len(args)+2)
	// not sure if that should be without path,  but having the path may expose private data?
	nargs = append(nargs, filepath.Base(file))
	nargs = append(nargs, line)
	errs := 0
	notNil := 0
	for _, arg := range args {
		if e, ok := arg.(error); ok {
			errs++
			if e != nil {
				notNil++
			}
		}
		nargs = append(nargs, arg)
	}
	// if there are error types in the argumens and they are all nil
	// we return just nil
	if errs > 0 && notNil == 0 {
		return nil
	}
	return fmt.Errorf(format, nargs)
}
