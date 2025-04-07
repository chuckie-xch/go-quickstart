package test

import "testing"

func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res != 3 {
		t.Errorf("expect %d, actual %d", 3, res)
	}
}
