package maps

import "testing"

func TestGetValue(t *testing.T) {
	// test int value
	var mValInt map[string]int
	if GetValue(mValInt, "k", 1) != 1 {
		t.Fatalf("unexpected value")
	}
	// test string value
	mValStr := map[string]string{"k1": "v1"}
	if GetValue(mValStr, "k", "na") != "na" {
		t.Fatalf("unexpected value")
	}

	// test int key
	mKeyInt64 := map[int64]string{1: "v1"}
	if GetValue(mKeyInt64, 2, "na") != "na" {
		t.Fatalf("unexpected value")
	}
}
