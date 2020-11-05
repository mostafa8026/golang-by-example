package unit_test_test

import (
	"testing"

	unit_test "github.com/mostafa8026/golang-by-example/unit-test"
)

func TestAdd(t *testing.T) {
	if unit_test.Add(2, 2) != 4 {
		t.Error("test failed.")
	}
}
