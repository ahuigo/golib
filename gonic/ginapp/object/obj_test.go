package object

import (
	"encoding/json"
	"testing"
)

func TestConvertMapBytes(t *testing.T) {
	mm := map[string][]byte{
		"k1": []byte("v1"),
		"k2": []byte("v2"),
	}
	obj := ConvertObjectByte2String(mm)
	if out, err := json.Marshal(obj); err != nil {
		t.Fatal(err)
	} else {
		expectedOut := `{"k1":"v1","k2":"v2"}`
		if string(out) != expectedOut {
			t.Fatalf("expected out:%v, unexpected out: %v", expectedOut, string(out))
		}
	}
}

func TestConvertOmitEmpty(t *testing.T) {
	type HistoryEvent struct {
		EventId *int64 `json:"eventId,omitempty"`
		TaskId  *int64 `json:"taskId,omitempty"`
	}
	i := int64(1)
	obj := HistoryEvent{
		EventId: &i,
	}
	if out, err := json.Marshal(ConvertObjectByte2String(obj)); err != nil {
		t.Fatal(err)
	} else {
		expectedOut := `{"eventId":1}`
		if string(out) != expectedOut {
			t.Fatalf("expected out:%v, unexpected out: %v", expectedOut, string(out))
		}
	}
}
