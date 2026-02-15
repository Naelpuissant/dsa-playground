package oll_test

import (
	oll "ds/ordered_linked_list"
	"reflect"
	"testing"
)

func TestOrderedLinkedList(t *testing.T) {
	list := oll.NewOrderedLinkedList()
	list.Insert(5, 1)
	list.Insert(3, 2)
	list.Insert(8, 3)
	list.Insert(1, 4)

	expectedKeys := []int{1, 3, 5, 8}
	actualKeys := list.Keys()
	if !reflect.DeepEqual(expectedKeys, actualKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}

	node := list.Find(3)
	if node != nil && node.Value != 2 {
		t.Errorf("Expected node value to be 2 got %v", node)
	}

	list.Delete(3)
	expectedKeys = []int{1, 5, 8}
	actualKeys = list.Keys()
	if !reflect.DeepEqual(actualKeys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}

	list.Delete(1)
	expectedKeys = []int{5, 8}
	actualKeys = list.Keys()
	if !reflect.DeepEqual(actualKeys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}

	list.Delete(8)
	expectedKeys = []int{5}
	actualKeys = list.Keys()
	if !reflect.DeepEqual(actualKeys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}

	node = list.Find(5)
	if node != nil && node.Value != 1 {
		t.Errorf("Expected node value to be 1 got %v", node)
	}

	list.Delete(5)
	expectedKeys = []int{}
	actualKeys = list.Keys()
	if !reflect.DeepEqual(actualKeys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}
}
