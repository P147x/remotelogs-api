package convert_test

import "testing"

func TestIntArrayConvertion(t *testing.T) {
	str := []int8{104, 101, 108, 108, 111}
	var res string
	for _, c := range str {
		res += string(c)
	}
	if res != "hello" {
		t.Errorf("Incorrect conversion, got: %s, want: %s.", res, "hello")
	}
}
