package arena

import (
	"errors"
	"sync/atomic"
)

var (
	ErrNotEnoughSpace = errors.New("Allocation failed, not enough space")
)

type Arena struct {
	buf []byte
	n   atomic.Uint64
}

func NewArena(size uint64) *Arena {
	return &Arena{
		buf: make([]byte, size),
		n:   atomic.Uint64{},
	}
}

func (a *Arena) Alloc(size uint64) (uint64, error) {
	for {
		// Using a loop here, but if we are sure that we discard the
		// arena once size reach buf limit, we can just use Add.
		n := a.n.Load()
		if n+size > uint64(len(a.buf)) {
			return 0, ErrNotEnoughSpace
		}
		if a.n.CompareAndSwap(n, n+size) {
			return n, nil
		}
	}
}

func (a *Arena) GetBytes(offset uint64, size uint64) []byte {
	return a.buf[offset:size]
}

func (a *Arena) Reset() {
	a.n.Store(0)
}
