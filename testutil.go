package ondotori

import (
	"runtime"
	"testing"
)

func testEquals(t *testing.T, expect interface{}, value interface{}) {
	if expect != value {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			t.Fatal("expect [", expect, "] but [", value, "]", file, line)
		} else {
			t.Fatal("expect [", expect, "] but [", value, "]")
		}
	}
}
