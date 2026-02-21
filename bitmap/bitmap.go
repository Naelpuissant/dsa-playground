package bm

import (
	"fmt"
	"strings"
)

func Uint64ToStr(i uint64) string {
	return fmt.Sprintf("%064b", i)
}

var (
	ErrBitMapIndexTooBig = fmt.Errorf("Value is too big, max value should be 63")
	ErrWrongBitMapSize   = fmt.Errorf("BitMap size should be a multiple of 64")
)

type BitMap struct {
	m []uint64
}

func NewBitMap(size int) (*BitMap, error) {
	if size%64 != 0 {
		return nil, ErrWrongBitMapSize
	}
	blocs := size / 64
	return &BitMap{m: make([]uint64, blocs)}, nil
}

func (bm *BitMap) getBlocForIndex(i int) (int, error) {
	if i >= len(bm.m)*64 {
		return -1, ErrBitMapIndexTooBig
	}

	bloc := i >> 6 // floor division by 64
	return bloc, nil
}

func (bm *BitMap) Set(i int) error {
	bloc, err := bm.getBlocForIndex(i)
	if err != nil {
		return err
	}

	shift := i % 64
	bm.m[len(bm.m)-1-bloc] |= 1 << shift

	return nil
}

func (bm *BitMap) Toggle(i int) error {
	bloc, err := bm.getBlocForIndex(i)
	if err != nil {
		return err
	}

	shift := i % 64
	bm.m[len(bm.m)-1-bloc] ^= 1 << shift

	return nil
}

func (bm *BitMap) IsSet(i int) bool {
	bloc, err := bm.getBlocForIndex(i)
	if err != nil {
		return false
	}

	shift := i % 64
	tmp := bm.m[len(bm.m)-1-bloc] | 1<<shift

	if tmp != bm.m[len(bm.m)-1-bloc] {
		return false
	}
	return true
}

func (bm *BitMap) UnSet(i int) error {
	if i >= len(bm.m)*64 {
		return ErrBitMapIndexTooBig
	}

	if bm.IsSet(i) {
		return bm.Toggle(i)
	}

	return nil
}

func (bm *BitMap) String() string {
	var str strings.Builder
	for _, bloc := range bm.m {
		str.WriteString(Uint64ToStr(bloc))
	}
	return str.String()
}

func (bm *BitMap) Size() int {
	return len(bm.m) * 64
}
