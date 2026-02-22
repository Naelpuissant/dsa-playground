package sl_test

import (
	sl "ds/skiplist"
	"math/rand"
	"reflect"
	"sync"
	"testing"

	"github.com/huandu/skiplist"
)

func TestSkiplist(t *testing.T) {
	list := sl.NewSkiplist()
	list.Insert(5, 1)
	list.Insert(3, 2)
	list.Insert(8, 3)
	list.Insert(1, 4)

	// Insert
	expectedKeys := []int{1, 3, 5, 8}
	actualKeys := list.Keys()
	if !reflect.DeepEqual(expectedKeys, actualKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, actualKeys)
	}

	// Search existing key
	node := list.Search(3)
	if node == nil || node.Value != 2 {
		t.Errorf("Expected to find key 3 with value 2, but got %v", node)
	}

	// Search non-existing key
	node = list.Search(10)
	if node != nil {
		t.Errorf("Expected to not find key 10, but got %v", node)
	}

	// Update value
	expectedLen := list.Length()
	list.Insert(3, 1337)
	if list.Length() != expectedLen {
		t.Errorf("Expected length %d, but got %d", expectedLen, list.Length())
	}

	// After update search
	node = list.Search(3)
	if node == nil || node.Value != 1337 {
		t.Errorf("Expected to find key 3 with value 1337, but got %v", node)
	}
}

func TestSkiplistInsertConcurrency(t *testing.T) {
	list := sl.NewSkiplist()

	wg := sync.WaitGroup{}
	wg.Add(1000)
	// Concurrent inserts
	for i := range 1000 {
		go func(i int) {
			list.Insert(i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Verify all keys are present
	for i := range 1000 {
		node := list.Search(i)
		if node == nil || node.Value != i {
			t.Errorf("Expected to find key %d with value %d, but got %v", i, i, node)
		}
	}
}

func BenchmarkSkiplistInsert(b *testing.B) {
	list := sl.NewSkiplist()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Insert(i, i)
	}
}

func BenchmarkHuanduSkiplistInsert(b *testing.B) {
	list := skiplist.New(skiplist.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}
}

func BenchmarkSkiplistSearch(b *testing.B) {
	list := sl.NewSkiplist()
	exepectedList := skiplist.New(skiplist.Int)

	len := 10000
	rands := make([]int, len)

	for i := range len {
		list.Insert(i, i)
		exepectedList.Set(i, i)
		rands[i] = rand.Intn(len - 1)
	}

	// Verify correctness before benchmarking
	for key := range len {
		node := list.Search(rands[key])
		expectedValue, _ := exepectedList.GetValue(rands[key])
		if node.Value != expectedValue {
			b.Fatalf("Incorect value %d for key %d, value should be %d", node.Value, rands[key], expectedValue)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Search(rands[i%len])
	}
}

func BenchmarkHuanduSkiplistSearch(b *testing.B) {
	list := skiplist.New(skiplist.Int)

	len := 10000
	rands := make([]int, len)
	for i := range len {
		list.Set(i, i)
		rands[i] = rand.Intn(len - 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(rands[i%len])
	}
}
