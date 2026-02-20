package bm_test

import (
	bm "ds/bitmap"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitmap := bm.NewBitMap()
	bitmap.Set(7)

	bmStr := bitmap.Str()
	c := bmStr[31-7] // stored right to left
	if c != '1' {
		t.Fatalf("BitMap[7] should be 1, got %b", c)
	}

	bitmap.Set(7)
	if bitmap.Str() != bmStr {
		t.Fatalf("BitMap should be unchanged after setting same bit, got %s", bitmap.Str())
	}

	if !bitmap.IsSet(7) {
		t.Fatalf("BitMap[7] should be set")
	}

	bitmap.Set(0)
	if !bitmap.IsSet(0) {
		t.Fatalf("BitMap[0] should be set")
	}

	bitmap.Set(31)
	if !bitmap.IsSet(31) {
		t.Fatalf("BitMap[31] should be set")
	}

	bitmap.Toggle(7)
	if bitmap.IsSet(7) {
		t.Fatalf("BitMap[7] shouldn't be set")
	}
	bitmap.Toggle(7)
	if !bitmap.IsSet(7) {
		t.Fatalf("BitMap[7] should be set again")
	}

	bitmap.UnSet(7)
	if bitmap.IsSet(7) {
		t.Fatalf("BitMap[7] should be unset")
	}

	err := bitmap.Set(32)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when setting bit index 32, got %v", err)
	}

	err = bitmap.UnSet(32)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when unsetting bit index 32, got %v", err)
	}

	err = bitmap.Toggle(32)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when toggling bit index 32, got %v", err)
	}
}
