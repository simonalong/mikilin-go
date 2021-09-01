package assert

import "testing"

func AssertTrue(t *testing.T, value bool) {
	if !value {
		t.Error("expect true, but actual is false")
	}
}

func AssertTrueErr(t *testing.T, value bool, errMsg string) {
	if !value {
		t.Errorf("expect true, but actual is false, error: %v", errMsg)
	}
}

func AssertFalse(t *testing.T, value bool) {
	if value {
		t.Error("expect false, but actual is true")
	}
}

func AssertFalseErr(t *testing.T, value bool, errMsg string) {
	if value {
		t.Errorf("expect false, but actual is true, error: %v", errMsg)
	}
}
