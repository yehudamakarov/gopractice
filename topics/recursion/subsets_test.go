package recursion

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2, 3})
	expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %#v but got %#v.", expected, got)
	}
}
