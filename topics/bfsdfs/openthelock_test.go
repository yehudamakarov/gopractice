package bfsdfs

import "testing"

func TestOpenLock(t *testing.T) {
	// got := OpenLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	// expected := 6

	got := OpenLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	expected := -1
	if got != expected {
		t.Errorf("Got: %d, expected: %d ", got, expected)
	}
}
