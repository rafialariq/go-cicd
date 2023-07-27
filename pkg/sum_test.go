package pkg

import "testing"

func TestSum(t *testing.T) {
	result := Sum(5)
	expected := 25
	if result != expected {
		t.Fail()
	}
}
