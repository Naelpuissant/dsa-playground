package sl_test

import (
	"ds/sl"
	"reflect"
	"testing"
)

func TestSkiplist(t *testing.T) {
	list := sl.NewSkiplist()
	list.Insert(5, 1)
	list.Insert(3, 2)
	list.Insert(8, 3)
	list.Insert(1, 4)

	expectedKeys := []int{1, 3, 5, 8}
	actualKeys := list.Keys()
	if !reflect.DeepEqual(expectedKeys, actualKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}
}
