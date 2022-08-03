package trace

import (
	"fmt"
    "testing"
	"runtime"
	"strings"
    "runtime/debug"
)

// Frame holds information about a single frame in the call stack.
type Frame struct {
	// Unique, package path-qualified name for the function of this call
	// frame.
	Function string

	// File and line number of our location in the frame.
	//
	// Note that the line number does not refer to where the function was
	// defined but where in the function the next call was made.
	File string
	Line int
}

func (f Frame) String() string {
	// This takes the following forms.
	//  (path/to/file.go)
	//  (path/to/file.go:42)
	//  path/to/package.MyFunction
	//  path/to/package.MyFunction (path/to/file.go)
	//  path/to/package.MyFunction (path/to/file.go:42)

	var sb strings.Builder
	sb.WriteString(f.Function)
	if len(f.File) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		fmt.Fprintf(&sb, "(%v", f.File)
		if f.Line > 0 {
			fmt.Fprintf(&sb, ":%d", f.Line)
		}
		sb.WriteRune(')')
	}

	if sb.Len() == 0 {
		return "unknown"
	}

	return sb.String()
}

func TestA(t *testing.T) {
    debug.PrintStack()
    CallerStack(t)
}



// CallerStack returns the call stack for the calling function, up to depth frames
// deep, skipping the provided number of frames, not including Callers itself.
//
// If zero, depth defaults to 8.
func CallerStack(t *testing.T) {
    depth := 8
    skip := 1

	pcs := make([]uintptr, depth)

	// +2 to skip this frame and runtime.Callers.
	n := runtime.Callers(skip, pcs)
	pcs = pcs[:n] // truncate to number of frames actually read

	frames := runtime.CallersFrames(pcs)
	for f, more := frames.Next(); more; f, more = frames.Next() {
		fmt.Println( f.File, f.Line, f.Function,)
	}
}
