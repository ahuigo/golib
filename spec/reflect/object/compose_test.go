package fn

import (
	"reflect"
	"testing"
)

type ClientTrace struct {
	ConnectStart func()
	// connectEnd   func() // reflect cannot access unexported fields
}

func composeTrace(trace, old *ClientTrace) {
	tv := reflect.ValueOf(trace).Elem()
	ov := reflect.ValueOf(old).Elem()
	for i := 0; i < tv.Type().NumField(); i++ {
		tf := tv.Field(i)
		of := ov.Field(i)

		//filter function
		if tf.Type().Kind() != reflect.Func {
			continue
		}

		//filter nil
		if of.IsNil() {
			continue
		}
		if tf.IsNil() {
			tf.Set(of)
			continue
		}

		// tfCopy := tf // stack overflow: tf is a pointer to tv.Field(i)
		// makeCopy: (Otherwise it creates a recursive call cycle )
		tfCopy := reflect.ValueOf(tf.Interface())

		// wrap
		newFunc := reflect.MakeFunc(tf.Type(), func(args []reflect.Value) []reflect.Value {
			tfCopy.Call(args)
			return of.Call(args)
		})
		tv.Field(i).Set(newFunc)
	}
}

func TestCompose(t *testing.T) {
	trace := &ClientTrace{
		ConnectStart: func() {
			println("connection started 1")
		},
	}
	traceOld := &ClientTrace{
		ConnectStart: func() {
			println("connection started 2.")
		},
	}
	_ = traceOld
	composeTrace(trace, traceOld)
	trace.ConnectStart()
}
