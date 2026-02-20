package bm

import "fmt"

func Uint32ToStr(i uint32) string {
	return fmt.Sprintf("%032b", i)
}

var (
	ErrBitMapIndexTooBig = fmt.Errorf("Value is too big, max value should be 31")
)

type BitMap struct {
	m uint32 // 32 byte long bitmap
}

func NewBitMap() *BitMap {
	return &BitMap{m: 0}
}

func (bm *BitMap) Set(i uint32) error {
	if i >= 32 {
		return ErrBitMapIndexTooBig
	}

	bm.m |= 1 << i

	return nil
}

func (bm *BitMap) UnSet(i uint32) error {
	if i >= 32 {
		return ErrBitMapIndexTooBig
	}

	if bm.IsSet(i) {
		return bm.Toggle(i)
	}

	return nil
}

func (bm *BitMap) Toggle(i uint32) error {
	if i >= 32 {
		return ErrBitMapIndexTooBig
	}

	bm.m ^= 1 << i

	return nil
}

func (bm *BitMap) IsSet(i uint32) bool {
	if i >= 32 {
		return false
	}

	tmp := bm.m | (uint32(1) << i)

	if tmp != bm.m {
		return false
	}
	return true
}

func (bm *BitMap) Str() string {
	return Uint32ToStr(bm.m)
}
