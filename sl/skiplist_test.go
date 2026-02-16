package sl_test

import (
	"ds/sl"
	"math/rand"
	"reflect"
	"testing"

	"github.com/huandu/skiplist"
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

	node := list.Search(3)
	if node == nil || node.Value != 2 {
		t.Errorf("Expected to find key 3 with value 2, but got %v", node)
	}
}

func BenchmarkSkiplistInsert(b *testing.B) {
	list := sl.NewSkiplist()
	for b.Loop() {
		list.Insert(b.N, b.N)
	}
}

func BenchmarkHuanduSkiplistInsert(b *testing.B) {
	list := skiplist.New(skiplist.Int)
	for b.Loop() {
		list.Set(b.N, b.N)
	}
}

func BenchmarkSkiplistSearch(b *testing.B) {
	list := sl.NewSkiplist()
	exepectedList := skiplist.New(skiplist.Int)

	len := 10000
	rands := []int{}

	for i := range len {
		list.Insert(i, i)
		exepectedList.Set(i, i)
		rands = append(rands, rand.Intn(len-1))
	}

	res := map[int]int{}
	for b.Loop() {
		node := list.Search(rands[b.N])
		res[rands[b.N]] = node.Value
	}

	for key, value := range res {
		expectedValue, _ := exepectedList.GetValue(key)
		if value != expectedValue {
			b.Fatalf("Incorect value %d for key %d, value should be %d", value, key, expectedValue)
		}
	}
}

func BenchmarkHuanduSkiplistSearch(b *testing.B) {
	list := skiplist.New(skiplist.Int)

	len := 10000
	rands := []int{}
	for i := range len {
		list.Set(i, i)
		rands = append(rands, rand.Intn(len-1))
	}

	for b.Loop() {
		list.Get(rands[b.N])
	}
}
