package bm_test

import (
	bm "ds/bitmap"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitmap, _ := bm.NewBitMap(256)
	bitmap.Set(7)

	_, err := bm.NewBitMap(255)
	if err != bm.ErrWrongBitMapSize {
		t.Fatalf("Expected error when creating BitMap of size 255, got %v", err)
	}

	bmStr := bitmap.String()
	c := bmStr[255-7] // stored right to left
	if c != '1' {
		t.Fatalf("BitMap[7] should be 1, got %b", c)
	}

	bitmap.Set(120)
	bmStr = bitmap.String()
	c = bmStr[255-120] // stored right to left
	if c != '1' {
		t.Fatalf("BitMap[120] should be 1, got %b", c)
	}

	bitmap.Set(7)
	if bitmap.String() != bmStr {
		t.Fatalf("BitMap should be unchanged after setting same bit, got %s", bitmap.String())
	}

	if !bitmap.IsSet(7) {
		t.Fatalf("BitMap[7] should be set")
	}

	bitmap.Set(0)
	if !bitmap.IsSet(0) {
		t.Fatalf("BitMap[0] should be set")
	}

	bitmap.Set(255)
	if !bitmap.IsSet(255) {
		t.Fatalf("BitMap[255] should be set")
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

	err = bitmap.Set(256)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when setting bit index 32, got %v", err)
	}

	err = bitmap.UnSet(256)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when unsetting bit index 32, got %v", err)
	}

	err = bitmap.Toggle(256)
	if err != bm.ErrBitMapIndexTooBig {
		t.Fatalf("Expected error when toggling bit index 32, got %v", err)
	}
}
