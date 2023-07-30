package lock

/**
A Pool is a set of temporary objects that may be individually saved and retrieved.

Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.

A Pool is safe for use by multiple goroutines simultaneously.
*/

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	// 协程间pool 复用
	b := bufPool.Get().(*bytes.Buffer)
	os.Stdout.WriteString("ori: " + b.String() + "\n")
	b.Reset()

	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	w.Write([]byte{'\n', '\n'})
	bufPool.Put(b)
}

func TestSyncPoll(t *testing.T) {
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=abc")
}
